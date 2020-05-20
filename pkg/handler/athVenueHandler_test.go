package handler

import (
	venuepb "backend/api/athVenue/v1"
	"backend/pkg/io"
	athVenue "backend/service/athVenue/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_CreateVenue(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthVenueService := athVenue.NewMockAthVenueService(ctrl)

	services := grpcServer{
		athVenueService: mockAthVenueService,
	}

	type args struct {
		ctx context.Context
		req *venuepb.CreateVenueRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *venuepb.CreateVenueReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
			},
			want: &venuepb.CreateVenueReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "OK with hours data",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueRequest{
					Hours: map[int32]string{1: "10:00-18:00"},
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
				mockAthVenueService.
					EXPECT().
					SaveHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &venuepb.CreateVenueReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreateVenue service error",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenue(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createvenue"), Success: false})
			},
			want: &venuepb.CreateVenueReply{
				Status: false,
				Error:  "failed to createvenue",
			},
			wantErr: false,
		},
		{
			name: "SaveHoursOfOperation service error",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueRequest{
					Hours: map[int32]string{1: "10:00-18:00"},
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
				mockAthVenueService.
					EXPECT().
					SaveHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to savehoursofoperation"), Success: false})
			},
			want: &venuepb.CreateVenueReply{
				Status: false,
				Error:  "failed to savehoursofoperation",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateVenue(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_EditVenue(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthVenueService := athVenue.NewMockAthVenueService(ctrl)

	services := grpcServer{
		athVenueService: mockAthVenueService,
	}

	type args struct {
		ctx context.Context
		req *venuepb.EditVenueRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *venuepb.GenericReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &venuepb.EditVenueRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					EditVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
			},
			want: &venuepb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "OK with Hours Data",
			args: args{
				ctx: ctx,
				req: &venuepb.EditVenueRequest{
					Hours:       map[int32]string{1: "10:00-18:00"},
					DeleteHours: []int32{2},
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					EditVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
				mockAthVenueService.
					EXPECT().
					EditHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthVenueService.
					EXPECT().
					DeleteHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &venuepb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "EditVenue service error",
			args: args{
				ctx: ctx,
				req: &venuepb.EditVenueRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					EditVenue(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to editvenue"), Success: false})
			},
			want: &venuepb.GenericReply{
				Status: false,
				Error:  "failed to editvenue",
			},
			wantErr: false,
		},
		{
			name: "EditHoursOfOperation service error",
			args: args{
				ctx: ctx,
				req: &venuepb.EditVenueRequest{
					Hours:       map[int32]string{1: "10:00-18:00"},
					DeleteHours: []int32{2},
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					EditVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
				mockAthVenueService.
					EXPECT().
					EditHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to edithoursofoperation"), Success: false})
			},
			want: &venuepb.GenericReply{
				Status: false,
				Error:  "failed to edithoursofoperation",
			},
			wantErr: false,
		},
		{
			name: "DeleteHoursOfOperation service error",
			args: args{
				ctx: ctx,
				req: &venuepb.EditVenueRequest{
					Hours:       map[int32]string{1: "10:00-18:00"},
					DeleteHours: []int32{2},
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					EditVenue(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_id": 1,
					}})
				mockAthVenueService.
					EXPECT().
					EditHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthVenueService.
					EXPECT().
					DeleteHoursOfOperation(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to deletehoursofoperation"), Success: false})
			},
			want: &venuepb.GenericReply{
				Status: false,
				Error:  "failed to deletehoursofoperation",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.EditVenue(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_CreateVenueHoliday(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthVenueService := athVenue.NewMockAthVenueService(ctrl)

	services := grpcServer{
		athVenueService: mockAthVenueService,
	}

	type args struct {
		ctx context.Context
		req *venuepb.CreateVenueHolidayRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *venuepb.CreateVenueHolidayReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueHolidayRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenueHoliday(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"venue_holiday_id": 1,
					}})
			},
			want: &venuepb.CreateVenueHolidayReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreateVenueHoliday service error",
			args: args{
				ctx: ctx,
				req: &venuepb.CreateVenueHolidayRequest{},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					CreateVenueHoliday(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createvenueholiday"), Success: false})
			},
			want: &venuepb.CreateVenueHolidayReply{
				Status: false,
				Error:  "failed to createvenueholiday",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateVenueHoliday(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_DeleteVenueHoliday(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthVenueService := athVenue.NewMockAthVenueService(ctrl)

	services := grpcServer{
		athVenueService: mockAthVenueService,
	}

	type args struct {
		ctx context.Context
		req *venuepb.DeleteVenueHolidayRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *venuepb.GenericReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &venuepb.DeleteVenueHolidayRequest{
					HolidayId: 1,
					VenueId:   1,
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					DeleteVenueHoliday(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &venuepb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "DeleteVenueHoliday service error",
			args: args{
				ctx: ctx,
				req: &venuepb.DeleteVenueHolidayRequest{
					HolidayId: 1,
					VenueId:   1,
				},
			},
			mock: func() {
				mockAthVenueService.
					EXPECT().
					DeleteVenueHoliday(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to deletevenueholiday"), Success: false})
			},
			want: &venuepb.GenericReply{
				Status: false,
				Error:  "failed to deletevenueholiday",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.DeleteVenueHoliday(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
