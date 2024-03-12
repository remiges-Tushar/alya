// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package batchsqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	FetchBatchRowsData(ctx context.Context, batch uuid.UUID) ([]FetchBatchRowsDataRow, error)
	FetchBlockOfRows(ctx context.Context, arg FetchBlockOfRowsParams) ([]FetchBlockOfRowsRow, error)
	GetBatchByID(ctx context.Context, id uuid.UUID) (Batch, error)
	GetBatchRowsByBatchIDSorted(ctx context.Context, batch uuid.UUID) ([]GetBatchRowsByBatchIDSortedRow, error)
	GetBatchStatus(ctx context.Context, id uuid.UUID) (StatusEnum, error)
	GetCompletedBatches(ctx context.Context) ([]uuid.UUID, error)
	GetPendingBatchRows(ctx context.Context, batch uuid.UUID) ([]GetPendingBatchRowsRow, error)
	InsertIntoBatchRows(ctx context.Context, arg InsertIntoBatchRowsParams) error
	InsertIntoBatches(ctx context.Context, arg InsertIntoBatchesParams) (uuid.UUID, error)
	UpdateBatchCounters(ctx context.Context, arg UpdateBatchCountersParams) error
	UpdateBatchOutputFiles(ctx context.Context, arg UpdateBatchOutputFilesParams) error
	UpdateBatchRowsBatchJob(ctx context.Context, arg UpdateBatchRowsBatchJobParams) error
	UpdateBatchRowsSlowQuery(ctx context.Context, arg UpdateBatchRowsSlowQueryParams) error
	UpdateBatchRowsStatus(ctx context.Context, arg UpdateBatchRowsStatusParams) error
	UpdateBatchSummary(ctx context.Context, arg UpdateBatchSummaryParams) error
}

var _ Querier = (*Queries)(nil)
