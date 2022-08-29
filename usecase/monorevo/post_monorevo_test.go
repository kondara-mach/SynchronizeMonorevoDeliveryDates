package monorevo_test

import (
	"SynchronizeMonorevoDeliveryDates/domain/monorevo"
	"SynchronizeMonorevoDeliveryDates/domain/monorevo/mock_monorevo"
	local "SynchronizeMonorevoDeliveryDates/usecase/monorevo"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestPropositionPostingUseCase_Execute(t *testing.T) {
	// logger生成
	logger, _ := zap.NewDevelopment()

	// PostRange戻り値
	mock_results := []monorevo.UpdatedProposition{
		*monorevo.TestUpdatedPropositionCreate(),
	}

	// モックコントローラーの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ものレボDIオブジェクト生成
	mock_poster := mock_monorevo.NewMockMonorevoPoster(ctrl)
	// EXPECTはctrl#Finishが呼び出される前に FetchAllを呼び出さなければエラーになる
	mock_poster.EXPECT().PostRange(gomock.Any()).Return(mock_results, nil)

	// UseCase戻り値
	results := []local.PostedPropositionDto{}
	for _, v := range mock_results {
		results = append(results,
			local.PostedPropositionDto{
				WorkedNumber:        v.WorkedNumber,
				Det:                 v.Det,
				Successful:          v.Successful,
				DeliveryDate:        v.DeliveryDate,
				UpdatedDeliveryDate: v.UpdatedDeliveryDate,
			},
		)
	}

	type args struct {
		p []local.PostingPropositionPram
	}
	tests := []struct {
		name    string
		m       *local.PropositionPostingUseCase
		args    args
		want    []local.PostedPropositionDto
		wantErr bool
	}{
		{
			name: "正常系_UseCaseを実行するとモックが実行されること",
			m: local.NewPropositionPostingUseCase(
				logger.Sugar(),
				mock_poster,
			),
			args: args{
				p: []local.PostingPropositionPram{
					{
						WorkedNumber:        "99A-1234",
						Det:                 "1",
						DeliveryDate:        time.Now(),
						UpdatedDeliveryDate: time.Now(),
					},
				},
			},
			want:    results,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Execute(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PropositionPostingUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropositionPostingUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
