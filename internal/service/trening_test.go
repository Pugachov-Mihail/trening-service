package service

import (
	"context"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
	"trening/internal/mocks"
	trening_v1 "trening/protos/gen/trening.v1"
)

func TestTreningService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrening := mocks.NewMockTreningStorage(ctrl)

	ctx := context.Background()
	log := slog.Default()

	mockTrening.EXPECT().
		TreningListSourse(ctx, int32(100), int32(0), gomock.Any()).
		Return([]trening_v1.GetTreningList{}, nil).
		Times(2)

	trening := Trening{Store: mockTrening, TStore: mockTrening, Log: log}

	result, err := trening.TreningsListService(ctx, 100, 0)

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 0 {
		t.Errorf("len(result) = %d, want %d", len(result), 0)
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
