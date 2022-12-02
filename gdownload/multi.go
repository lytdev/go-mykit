package gdownload

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"

	hos "github.com/lytdev/go-mykit/helpers/os"
	"github.com/panjf2000/ants/v2"
)

var pool *ants.PoolWithFunc

const (
	suffixDb = ".db"
	suffixDp = ".dp"
)

type invokeParams struct {
	ctx       context.Context
	completed chan *part
	failed    chan error
	buf       []byte
	dst       string
	part      *part
	listener  ProgressListener
	fileSize  int64
	reader    FileReader
}

// 文件的分片信息
type part struct {
	//索引
	id int `json:"id"`
	// 分片文件的起始位置
	offset int64 `json:"offset"`
	//分片文件的大小
	size int64 `json:"size"`
	//是否已下载完成
	isCompleted bool `json:"is_completed"`
}

// MultiDownload 进行文件分片下载
func (d *Downloader) MultiDownload(ctx context.Context, dst string, reader FileReader, listener ProgressListener) (err error) {
	defer ants.Release()
	//将下载任务放入线程池
	pool, _ = ants.NewPoolWithFunc(d.Workers, func(i interface{}) {
		params := i.(*invokeParams)
		downloadPartToFile(params)
	})
	db := dbPath(dst)
	var partInfoArr []*part
	var fileSize int64 = 0
	if hos.FileExists(db) {
		partInfoArr, err = loadDb(db)
	} else {
		fileSize, err = reader.GetFileSize()
		if err != nil {
			return err
		}
		// 下载文件的分片信息
		partInfoArr, err = d.initPartInfo(fileSize)
	}
	//已完成的数量
	completedCount := 0
	for _, part := range partInfoArr {
		if part.isCompleted {
			completedCount++
		}
	}
	//分片数
	partCount := len(partInfoArr)
	//已经完成的分片
	completed := make(chan *part, partCount)
	failed := make(chan error)
	listener.ProgressChanged(&ProgressEvent{
		ConsumedBytes: 0,
		TotalBytes:    fileSize,
		RwBytes:       0,
		EventType:     TransferStartedEvent,
	})
	for _, part := range partInfoArr {
		if !part.isCompleted {
			err = pool.Invoke(&invokeParams{
				ctx:       ctx,
				completed: completed,
				failed:    failed,
				buf:       make([]byte, d.BufSize),
				dst:       dpPath(part.id, dst),
				part:      part,
				fileSize:  fileSize,
				reader:    reader,
			})
			if err != nil {
				return err
			}
		}
	}
	var wm int64 = 0
	for completedCount < partCount {
		select {
		case rp := <-completed:
			completedCount++
			rp.isCompleted = true
			marshal, err := json.Marshal(partInfoArr)
			if err != nil {
				return err
			}
			os.WriteFile(db, marshal, os.ModePerm)
			wm += rp.size
			listener.ProgressChanged(&ProgressEvent{
				ConsumedBytes: wm,
				TotalBytes:    fileSize,
				RwBytes:       wm,
				EventType:     TransferDataEvent,
			})
		case err = <-failed:
			return err
		}

		if completedCount >= partCount {
			break
		}
	}
	return merge(dst, partInfoArr, fileSize, listener)
}

func (d *Downloader) Tune(size int) {
	pool.Tune(size)
}

// dbPath 拼接文件保存路径
func dbPath(dst string) string {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	return filepath.Join(dstDir, fmt.Sprintf("%s%s", dstBase, suffixDb))
}

// merge 分片文件的合并
func merge(dst string, parts []*part, fileSize int64, listener ProgressListener) error {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	newFilename, err := hos.NewFilename(dst, 10, nil)
	if err != nil {
		return err
	}

	fs, err := os.Create(newFilename)
	defer fs.Close()
	if err != nil {
		return err
	}
	var offset int64 = 0
	for i := 0; i < len(parts); i++ {
		pf := filepath.Join(dstDir, fmt.Sprintf("%s%s%d", dstBase, suffixDp, i))
		buf, err := os.ReadFile(pf)
		if err != nil {
			return err
		}
		at, err := fs.WriteAt(buf, offset)
		if err != nil {
			return err
		}
		offset += int64(at)
	}

	os.Remove(dbPath(dst))
	for i := 0; i < len(parts); i++ {
		os.Remove(dpPath(i, dst))
	}
	listener.ProgressChanged(&ProgressEvent{
		ConsumedBytes: fileSize,
		TotalBytes:    fileSize,
		RwBytes:       fileSize,
		EventType:     TransferCompletedEvent,
	})
	return nil
}

// initPartInfo 获取文件的分片详情
func (d *Downloader) initPartInfo(fileSize int64) (parts []*part, err error) {
	//整个下载的文件一共分多少片
	count := int(math.Ceil(float64(fileSize) / float64(d.PartSize)))
	var offset int64 = 0
	remain := fileSize
	parts = make([]*part, count)
	//循环获取每一片的信息
	for i := 0; i < count; i++ {
		ps := d.PartSize
		if remain < d.PartSize {
			ps = remain
		}
		remain -= d.PartSize
		parts[i] = &part{
			id:          i,
			offset:      offset,
			size:        ps,
			isCompleted: false,
		}
		offset += d.PartSize
	}
	return parts, nil
}

func dpPath(id int, dst string) string {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	return filepath.Join(dstDir, fmt.Sprintf("%s%s%d", dstBase, suffixDp, id))
}

// 加载已经下载的文件分片信息
func loadDb(dbPath string) (parts []*part, err error) {
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &parts)
	if err != nil {
		return parts, err
	}
	return parts, nil
}

// downloadPartToFile 下载分片至本地文件
func downloadPartToFile(params *invokeParams) (int64, error) {
	fs, err := os.Create(params.dst)
	defer fs.Close()
	if err != nil {
		return 0, err
	}
	return downloadPartToWriter(fs, params)
}

// downloadPartToWriter 下载分片文件至流
func downloadPartToWriter(writer io.WriterAt, params *invokeParams) (int64, error) {
	var err error
	select {
	case <-params.ctx.Done():
		err = errors.New("context done")
		params.failed <- err
		return -1, err
	default:
	}

	if params.part.size < 1 {
		err = errors.New("part size error")
		params.failed <- err
		return -1, err
	}

	if params.buf == nil {
		params.buf = make([]byte, 1024*200)
	}

	var wn int64 = 0

	body, err := params.reader.OpenRange(params.part.offset, params.part.size)
	if err != nil {
		return 0, err
	}
	defer body.Close()
	for {
		select {
		case <-params.ctx.Done():
			params.failed <- err
			return wn, errors.New("context done")
		default:
		}
		var n int
		n, err = body.Read(params.buf)
		if err != nil && err != io.EOF {
			params.failed <- err
			return wn, err
		} else {

			writer.WriteAt(params.buf[:n], wn)
			wn += int64(n)
			if err == io.EOF {
				break
			}

		}
	}
	params.completed <- params.part
	return wn, nil
}
