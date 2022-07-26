package jobbook_fetch_case

import (
	"SynchronizeMonorevoDeliveryDates/domain/orderdb"
	"time"

	"go.uber.org/zap"
)

type JobBookDto struct {
	WorkedNumber string
	DeliveryDate time.Time
}

type Executor interface {
	Execute() ([]JobBookDto, error)
}

type JobBookFetchingUseCase struct {
	sugar          *zap.SugaredLogger
	jobBookFetcher orderdb.JobBookFetcher
}

func NewJobBookFetchingUseCase(
	sugar *zap.SugaredLogger,
	jobBookFetcher orderdb.JobBookFetcher,
) *JobBookFetchingUseCase {
	return &JobBookFetchingUseCase{
		sugar:          sugar,
		jobBookFetcher: jobBookFetcher,
	}
}

func (m *JobBookFetchingUseCase) Execute() ([]JobBookDto, error) {
	job, err := m.jobBookFetcher.FetchAll()
	if err != nil {
		m.sugar.Fatal("受注管理DBから作業台帳を取得できませんでした", err)
	}

	// 詰め替え
	dto := []JobBookDto{}
	for _, v := range job {
		dto = append(
			dto,
			JobBookDto{
				WorkedNumber: v.WorkedNumber,
				DeliveryDate: v.DeliveryDate,
			},
		)
	}
	return dto, nil
}
