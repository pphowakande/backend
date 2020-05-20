package handler

import (
	facilitypb "backend/api/athFacility/v1"
	"backend/pkg/io"
	athFacility "backend/service/athFacility/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_CreateFacility(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthFacilityService := athFacility.NewMockAthFacilityService(ctrl)

	services := grpcServer{
		athFacilityService: mockAthFacilityService,
	}

	type args struct {
		ctx context.Context
		req *facilitypb.CreateFacilityRequest
	}

	weekData := facilitypb.WeekData{
		Weekdays: map[string]*facilitypb.Weekslots{
			"A": &facilitypb.Weekslots{
				Price: 0.00,
				Slot:  "10:00-19:00",
			},
		},
		Weekends: map[string]*facilitypb.Weekslots{
			"B": &facilitypb.Weekslots{
				Price: 0.00,
				Slot:  "09:00-19:00",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *facilitypb.CreateFacilityReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &facilitypb.CreateFacilityRequest{
					WeekData: &weekData,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					CreateFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"user_id":     1,
						"facility_id": 1,
					}})
				mockAthFacilityService.
					EXPECT().
					CreateFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}}).
					AnyTimes()
			},
			want: &facilitypb.CreateFacilityReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreateFacility service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.CreateFacilityRequest{
					WeekData: &weekData,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					CreateFacility(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createfacility"), Success: false})
			},
			want: &facilitypb.CreateFacilityReply{
				Status: false,
				Error:  "failed to createfacility",
			},
			wantErr: false,
		},
		{
			name: "CreateFacilitySlots service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.CreateFacilityRequest{
					WeekData: &weekData,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					CreateFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"user_id":     1,
						"facility_id": 1,
					}})
				mockAthFacilityService.
					EXPECT().
					CreateFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createfacilityslots"), Success: false}).
					AnyTimes()
			},
			want: &facilitypb.CreateFacilityReply{
				Status: false,
				Error:  "failed to createfacilityslots",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateFacility(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_EditFacility(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthFacilityService := athFacility.NewMockAthFacilityService(ctrl)

	services := grpcServer{
		athFacilityService: mockAthFacilityService,
	}

	type args struct {
		ctx context.Context
		req *facilitypb.EditFacilityRequest
	}

	weekDataEdit := facilitypb.WeekDataEdit{
		Weekdays: map[string]float32{
			"A": 1.0,
		},
		Weekends: map[string]float32{
			"B": 2.0,
		},
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *facilitypb.GenericReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &facilitypb.EditFacilityRequest{
					WeekData: &weekDataEdit,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					EditFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthFacilityService.
					EXPECT().
					EditFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}}).
					AnyTimes()
			},
			want: &facilitypb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "EditFacility service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.EditFacilityRequest{
					WeekData: &weekDataEdit,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					EditFacility(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to editfacility"), Success: false})
			},
			want: &facilitypb.GenericReply{
				Status: false,
				Error:  "failed to editfacility",
			},
			wantErr: false,
		},
		{
			name: "EditFacilitySlots service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.EditFacilityRequest{
					WeekData: &weekDataEdit,
				},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					EditFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthFacilityService.
					EXPECT().
					EditFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to editfacilityslots"), Success: false}).
					AnyTimes()
			},
			want: &facilitypb.GenericReply{
				Status: false,
				Error:  "failed to editfacilityslots",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.EditFacility(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_BookFacility(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthFacilityService := athFacility.NewMockAthFacilityService(ctrl)

	services := grpcServer{
		athFacilityService: mockAthFacilityService,
	}

	type args struct {
		ctx context.Context
		req *facilitypb.BookFacilityRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *facilitypb.BookFacilityReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &facilitypb.BookFacilityRequest{},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					BookFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"facility_booking_id": 1,
						"facility_booking_no": "test_no",
					}})
				mockAthFacilityService.
					EXPECT().
					BookFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &facilitypb.BookFacilityReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "BookFacility service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.BookFacilityRequest{},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					BookFacility(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to bookfacility"), Success: false})
			},
			want: &facilitypb.BookFacilityReply{
				Status: false,
				Error:  "failed to bookfacility",
			},
			wantErr: false,
		},
		{
			name: "BookFacilitySlots service error",
			args: args{
				ctx: ctx,
				req: &facilitypb.BookFacilityRequest{},
			},
			mock: func() {
				mockAthFacilityService.
					EXPECT().
					BookFacility(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"facility_booking_id": 1,
						"facility_booking_no": "test_no",
					}})
				mockAthFacilityService.
					EXPECT().
					BookFacilitySlots(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to bookfacilityslots"), Success: false})
			},
			want: &facilitypb.BookFacilityReply{
				Status: false,
				Error:  "failed to bookfacilityslots",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.BookFacility(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
