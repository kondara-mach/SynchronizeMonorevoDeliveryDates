package appsetting_test

import (
	local "SynchronizeMonorevoDeliveryDates/infrastructure/appsetting"
	"SynchronizeMonorevoDeliveryDates/usecase/appsetting_obtain_case"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func createDummyAppSetting(dummySetting *local.AppSetting) string {
	// テスト実行フォルダ取得
	exeFile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Join(filepath.Dir(exeFile), "TestData")
	if err := os.MkdirAll(exePath, 0777); err != nil {
		panic(err)
	}
	dummy := filepath.Join(exePath, "appsetting.json")

	// json書き込み
	f, _ := os.Create(dummy)
	if err := json.NewEncoder(f).Encode(dummySetting); err != nil {
		panic(err)
	}

	return dummy
}

func TestLoadableSetting_Load(t *testing.T) {
	// 仮の設定値
	dummySetting := local.TestAppSettingCreate()
	dummyPath := createDummyAppSetting(dummySetting)

	logger, _ := zap.NewDevelopment()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		l       *local.LoadableSetting
		args    args
		want    *appsetting_obtain_case.AppSettingDto
		wantErr bool
	}{
		{
			name: "正常系_設定値が取得できること",
			l:    local.NewLoadableSetting(logger.Sugar()),
			args: args{
				path: dummyPath,
			},
			want: &appsetting_obtain_case.AppSettingDto{
				SandboxMode: appsetting_obtain_case.SandboxModeDto{
					Monorevo: dummySetting.SandboxMode.Monorevo,
					SendGrid: dummySetting.SandboxMode.SendGrid,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Load(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadableSetting.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadableSetting.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
