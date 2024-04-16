// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/google/uuid"
	"github.com/remiges-tech/alya/batch/pg/batchsqlc"
	"sync"
)

// Ensure, that QuerierMock does implement batchsqlc.Querier.
// If this is not the case, regenerate this file with moq.
var _ batchsqlc.Querier = &QuerierMock{}

// QuerierMock is a mock implementation of batchsqlc.Querier.
//
//	func TestSomethingThatUsesQuerier(t *testing.T) {
//
//		// make and configure a mocked batchsqlc.Querier
//		mockedQuerier := &QuerierMock{
//			CountBatchRowsByBatchIDAndStatusFunc: func(ctx context.Context, arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams) (int64, error) {
//				panic("mock out the CountBatchRowsByBatchIDAndStatus method")
//			},
//			FetchBatchRowsDataFunc: func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.FetchBatchRowsDataRow, error) {
//				panic("mock out the FetchBatchRowsData method")
//			},
//			FetchBlockOfRowsFunc: func(ctx context.Context, arg batchsqlc.FetchBlockOfRowsParams) ([]batchsqlc.FetchBlockOfRowsRow, error) {
//				panic("mock out the FetchBlockOfRows method")
//			},
//			GetBatchByIDFunc: func(ctx context.Context, id uuid.UUID) (batchsqlc.Batch, error) {
//				panic("mock out the GetBatchByID method")
//			},
//			GetBatchRowsByBatchIDFunc: func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.Batchrow, error) {
//				panic("mock out the GetBatchRowsByBatchID method")
//			},
//			GetBatchRowsByBatchIDSortedFunc: func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetBatchRowsByBatchIDSortedRow, error) {
//				panic("mock out the GetBatchRowsByBatchIDSorted method")
//			},
//			GetBatchRowsCountFunc: func(ctx context.Context, batch uuid.UUID) (int64, error) {
//				panic("mock out the GetBatchRowsCount method")
//			},
//			GetBatchStatusFunc: func(ctx context.Context, id uuid.UUID) (batchsqlc.StatusEnum, error) {
//				panic("mock out the GetBatchStatus method")
//			},
//			GetCompletedBatchesFunc: func(ctx context.Context) ([]uuid.UUID, error) {
//				panic("mock out the GetCompletedBatches method")
//			},
//			GetPendingBatchRowsFunc: func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetPendingBatchRowsRow, error) {
//				panic("mock out the GetPendingBatchRows method")
//			},
//			GetProcessedBatchRowsByBatchIDSortedFunc: func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetProcessedBatchRowsByBatchIDSortedRow, error) {
//				panic("mock out the GetProcessedBatchRowsByBatchIDSorted method")
//			},
//			InsertIntoBatchRowsFunc: func(ctx context.Context, arg batchsqlc.InsertIntoBatchRowsParams) error {
//				panic("mock out the InsertIntoBatchRows method")
//			},
//			InsertIntoBatchesFunc: func(ctx context.Context, arg batchsqlc.InsertIntoBatchesParams) (uuid.UUID, error) {
//				panic("mock out the InsertIntoBatches method")
//			},
//			UpdateBatchCountersFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchCountersParams) error {
//				panic("mock out the UpdateBatchCounters method")
//			},
//			UpdateBatchOutputFilesFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchOutputFilesParams) error {
//				panic("mock out the UpdateBatchOutputFiles method")
//			},
//			UpdateBatchRowsBatchJobFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchRowsBatchJobParams) error {
//				panic("mock out the UpdateBatchRowsBatchJob method")
//			},
//			UpdateBatchRowsSlowQueryFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchRowsSlowQueryParams) error {
//				panic("mock out the UpdateBatchRowsSlowQuery method")
//			},
//			UpdateBatchRowsStatusFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchRowsStatusParams) error {
//				panic("mock out the UpdateBatchRowsStatus method")
//			},
//			UpdateBatchStatusFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchStatusParams) error {
//				panic("mock out the UpdateBatchStatus method")
//			},
//			UpdateBatchSummaryFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchSummaryParams) error {
//				panic("mock out the UpdateBatchSummary method")
//			},
//			UpdateBatchSummaryOnAbortFunc: func(ctx context.Context, arg batchsqlc.UpdateBatchSummaryOnAbortParams) error {
//				panic("mock out the UpdateBatchSummaryOnAbort method")
//			},
//		}
//
//		// use mockedQuerier in code that requires batchsqlc.Querier
//		// and then make assertions.
//
//	}
type QuerierMock struct {
	// CountBatchRowsByBatchIDAndStatusFunc mocks the CountBatchRowsByBatchIDAndStatus method.
	CountBatchRowsByBatchIDAndStatusFunc func(ctx context.Context, arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams) (int64, error)

	// FetchBatchRowsDataFunc mocks the FetchBatchRowsData method.
	FetchBatchRowsDataFunc func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.FetchBatchRowsDataRow, error)

	// FetchBlockOfRowsFunc mocks the FetchBlockOfRows method.
	FetchBlockOfRowsFunc func(ctx context.Context, arg batchsqlc.FetchBlockOfRowsParams) ([]batchsqlc.FetchBlockOfRowsRow, error)

	// GetBatchByIDFunc mocks the GetBatchByID method.
	GetBatchByIDFunc func(ctx context.Context, id uuid.UUID) (batchsqlc.Batch, error)

	// GetBatchRowsByBatchIDFunc mocks the GetBatchRowsByBatchID method.
	GetBatchRowsByBatchIDFunc func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.Batchrow, error)

	// GetBatchRowsByBatchIDSortedFunc mocks the GetBatchRowsByBatchIDSorted method.
	GetBatchRowsByBatchIDSortedFunc func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetBatchRowsByBatchIDSortedRow, error)

	// GetBatchRowsCountFunc mocks the GetBatchRowsCount method.
	GetBatchRowsCountFunc func(ctx context.Context, batch uuid.UUID) (int64, error)

	// GetBatchStatusFunc mocks the GetBatchStatus method.
	GetBatchStatusFunc func(ctx context.Context, id uuid.UUID) (batchsqlc.StatusEnum, error)

	// GetCompletedBatchesFunc mocks the GetCompletedBatches method.
	GetCompletedBatchesFunc func(ctx context.Context) ([]uuid.UUID, error)

	// GetPendingBatchRowsFunc mocks the GetPendingBatchRows method.
	GetPendingBatchRowsFunc func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetPendingBatchRowsRow, error)

	// GetProcessedBatchRowsByBatchIDSortedFunc mocks the GetProcessedBatchRowsByBatchIDSorted method.
	GetProcessedBatchRowsByBatchIDSortedFunc func(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetProcessedBatchRowsByBatchIDSortedRow, error)

	// InsertIntoBatchRowsFunc mocks the InsertIntoBatchRows method.
	InsertIntoBatchRowsFunc func(ctx context.Context, arg batchsqlc.InsertIntoBatchRowsParams) error

	// InsertIntoBatchesFunc mocks the InsertIntoBatches method.
	InsertIntoBatchesFunc func(ctx context.Context, arg batchsqlc.InsertIntoBatchesParams) (uuid.UUID, error)

	// UpdateBatchCountersFunc mocks the UpdateBatchCounters method.
	UpdateBatchCountersFunc func(ctx context.Context, arg batchsqlc.UpdateBatchCountersParams) error

	// UpdateBatchOutputFilesFunc mocks the UpdateBatchOutputFiles method.
	UpdateBatchOutputFilesFunc func(ctx context.Context, arg batchsqlc.UpdateBatchOutputFilesParams) error

	// UpdateBatchRowsBatchJobFunc mocks the UpdateBatchRowsBatchJob method.
	UpdateBatchRowsBatchJobFunc func(ctx context.Context, arg batchsqlc.UpdateBatchRowsBatchJobParams) error

	// UpdateBatchRowsSlowQueryFunc mocks the UpdateBatchRowsSlowQuery method.
	UpdateBatchRowsSlowQueryFunc func(ctx context.Context, arg batchsqlc.UpdateBatchRowsSlowQueryParams) error

	// UpdateBatchRowsStatusFunc mocks the UpdateBatchRowsStatus method.
	UpdateBatchRowsStatusFunc func(ctx context.Context, arg batchsqlc.UpdateBatchRowsStatusParams) error

	// UpdateBatchStatusFunc mocks the UpdateBatchStatus method.
	UpdateBatchStatusFunc func(ctx context.Context, arg batchsqlc.UpdateBatchStatusParams) error

	// UpdateBatchSummaryFunc mocks the UpdateBatchSummary method.
	UpdateBatchSummaryFunc func(ctx context.Context, arg batchsqlc.UpdateBatchSummaryParams) error

	// UpdateBatchSummaryOnAbortFunc mocks the UpdateBatchSummaryOnAbort method.
	UpdateBatchSummaryOnAbortFunc func(ctx context.Context, arg batchsqlc.UpdateBatchSummaryOnAbortParams) error

	// calls tracks calls to the methods.
	calls struct {
		// CountBatchRowsByBatchIDAndStatus holds details about calls to the CountBatchRowsByBatchIDAndStatus method.
		CountBatchRowsByBatchIDAndStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams
		}
		// FetchBatchRowsData holds details about calls to the FetchBatchRowsData method.
		FetchBatchRowsData []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// FetchBlockOfRows holds details about calls to the FetchBlockOfRows method.
		FetchBlockOfRows []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.FetchBlockOfRowsParams
		}
		// GetBatchByID holds details about calls to the GetBatchByID method.
		GetBatchByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetBatchRowsByBatchID holds details about calls to the GetBatchRowsByBatchID method.
		GetBatchRowsByBatchID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// GetBatchRowsByBatchIDSorted holds details about calls to the GetBatchRowsByBatchIDSorted method.
		GetBatchRowsByBatchIDSorted []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// GetBatchRowsCount holds details about calls to the GetBatchRowsCount method.
		GetBatchRowsCount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// GetBatchStatus holds details about calls to the GetBatchStatus method.
		GetBatchStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetCompletedBatches holds details about calls to the GetCompletedBatches method.
		GetCompletedBatches []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetPendingBatchRows holds details about calls to the GetPendingBatchRows method.
		GetPendingBatchRows []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// GetProcessedBatchRowsByBatchIDSorted holds details about calls to the GetProcessedBatchRowsByBatchIDSorted method.
		GetProcessedBatchRowsByBatchIDSorted []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch uuid.UUID
		}
		// InsertIntoBatchRows holds details about calls to the InsertIntoBatchRows method.
		InsertIntoBatchRows []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.InsertIntoBatchRowsParams
		}
		// InsertIntoBatches holds details about calls to the InsertIntoBatches method.
		InsertIntoBatches []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.InsertIntoBatchesParams
		}
		// UpdateBatchCounters holds details about calls to the UpdateBatchCounters method.
		UpdateBatchCounters []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchCountersParams
		}
		// UpdateBatchOutputFiles holds details about calls to the UpdateBatchOutputFiles method.
		UpdateBatchOutputFiles []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchOutputFilesParams
		}
		// UpdateBatchRowsBatchJob holds details about calls to the UpdateBatchRowsBatchJob method.
		UpdateBatchRowsBatchJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchRowsBatchJobParams
		}
		// UpdateBatchRowsSlowQuery holds details about calls to the UpdateBatchRowsSlowQuery method.
		UpdateBatchRowsSlowQuery []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchRowsSlowQueryParams
		}
		// UpdateBatchRowsStatus holds details about calls to the UpdateBatchRowsStatus method.
		UpdateBatchRowsStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchRowsStatusParams
		}
		// UpdateBatchStatus holds details about calls to the UpdateBatchStatus method.
		UpdateBatchStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchStatusParams
		}
		// UpdateBatchSummary holds details about calls to the UpdateBatchSummary method.
		UpdateBatchSummary []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchSummaryParams
		}
		// UpdateBatchSummaryOnAbort holds details about calls to the UpdateBatchSummaryOnAbort method.
		UpdateBatchSummaryOnAbort []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg batchsqlc.UpdateBatchSummaryOnAbortParams
		}
	}
	lockCountBatchRowsByBatchIDAndStatus     sync.RWMutex
	lockFetchBatchRowsData                   sync.RWMutex
	lockFetchBlockOfRows                     sync.RWMutex
	lockGetBatchByID                         sync.RWMutex
	lockGetBatchRowsByBatchID                sync.RWMutex
	lockGetBatchRowsByBatchIDSorted          sync.RWMutex
	lockGetBatchRowsCount                    sync.RWMutex
	lockGetBatchStatus                       sync.RWMutex
	lockGetCompletedBatches                  sync.RWMutex
	lockGetPendingBatchRows                  sync.RWMutex
	lockGetProcessedBatchRowsByBatchIDSorted sync.RWMutex
	lockInsertIntoBatchRows                  sync.RWMutex
	lockInsertIntoBatches                    sync.RWMutex
	lockUpdateBatchCounters                  sync.RWMutex
	lockUpdateBatchOutputFiles               sync.RWMutex
	lockUpdateBatchRowsBatchJob              sync.RWMutex
	lockUpdateBatchRowsSlowQuery             sync.RWMutex
	lockUpdateBatchRowsStatus                sync.RWMutex
	lockUpdateBatchStatus                    sync.RWMutex
	lockUpdateBatchSummary                   sync.RWMutex
	lockUpdateBatchSummaryOnAbort            sync.RWMutex
}

// CountBatchRowsByBatchIDAndStatus calls CountBatchRowsByBatchIDAndStatusFunc.
func (mock *QuerierMock) CountBatchRowsByBatchIDAndStatus(ctx context.Context, arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams) (int64, error) {
	if mock.CountBatchRowsByBatchIDAndStatusFunc == nil {
		panic("QuerierMock.CountBatchRowsByBatchIDAndStatusFunc: method is nil but Querier.CountBatchRowsByBatchIDAndStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockCountBatchRowsByBatchIDAndStatus.Lock()
	mock.calls.CountBatchRowsByBatchIDAndStatus = append(mock.calls.CountBatchRowsByBatchIDAndStatus, callInfo)
	mock.lockCountBatchRowsByBatchIDAndStatus.Unlock()
	return mock.CountBatchRowsByBatchIDAndStatusFunc(ctx, arg)
}

// CountBatchRowsByBatchIDAndStatusCalls gets all the calls that were made to CountBatchRowsByBatchIDAndStatus.
// Check the length with:
//
//	len(mockedQuerier.CountBatchRowsByBatchIDAndStatusCalls())
func (mock *QuerierMock) CountBatchRowsByBatchIDAndStatusCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.CountBatchRowsByBatchIDAndStatusParams
	}
	mock.lockCountBatchRowsByBatchIDAndStatus.RLock()
	calls = mock.calls.CountBatchRowsByBatchIDAndStatus
	mock.lockCountBatchRowsByBatchIDAndStatus.RUnlock()
	return calls
}

// FetchBatchRowsData calls FetchBatchRowsDataFunc.
func (mock *QuerierMock) FetchBatchRowsData(ctx context.Context, batch uuid.UUID) ([]batchsqlc.FetchBatchRowsDataRow, error) {
	if mock.FetchBatchRowsDataFunc == nil {
		panic("QuerierMock.FetchBatchRowsDataFunc: method is nil but Querier.FetchBatchRowsData was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockFetchBatchRowsData.Lock()
	mock.calls.FetchBatchRowsData = append(mock.calls.FetchBatchRowsData, callInfo)
	mock.lockFetchBatchRowsData.Unlock()
	return mock.FetchBatchRowsDataFunc(ctx, batch)
}

// FetchBatchRowsDataCalls gets all the calls that were made to FetchBatchRowsData.
// Check the length with:
//
//	len(mockedQuerier.FetchBatchRowsDataCalls())
func (mock *QuerierMock) FetchBatchRowsDataCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockFetchBatchRowsData.RLock()
	calls = mock.calls.FetchBatchRowsData
	mock.lockFetchBatchRowsData.RUnlock()
	return calls
}

// FetchBlockOfRows calls FetchBlockOfRowsFunc.
func (mock *QuerierMock) FetchBlockOfRows(ctx context.Context, arg batchsqlc.FetchBlockOfRowsParams) ([]batchsqlc.FetchBlockOfRowsRow, error) {
	if mock.FetchBlockOfRowsFunc == nil {
		panic("QuerierMock.FetchBlockOfRowsFunc: method is nil but Querier.FetchBlockOfRows was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.FetchBlockOfRowsParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockFetchBlockOfRows.Lock()
	mock.calls.FetchBlockOfRows = append(mock.calls.FetchBlockOfRows, callInfo)
	mock.lockFetchBlockOfRows.Unlock()
	return mock.FetchBlockOfRowsFunc(ctx, arg)
}

// FetchBlockOfRowsCalls gets all the calls that were made to FetchBlockOfRows.
// Check the length with:
//
//	len(mockedQuerier.FetchBlockOfRowsCalls())
func (mock *QuerierMock) FetchBlockOfRowsCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.FetchBlockOfRowsParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.FetchBlockOfRowsParams
	}
	mock.lockFetchBlockOfRows.RLock()
	calls = mock.calls.FetchBlockOfRows
	mock.lockFetchBlockOfRows.RUnlock()
	return calls
}

// GetBatchByID calls GetBatchByIDFunc.
func (mock *QuerierMock) GetBatchByID(ctx context.Context, id uuid.UUID) (batchsqlc.Batch, error) {
	if mock.GetBatchByIDFunc == nil {
		panic("QuerierMock.GetBatchByIDFunc: method is nil but Querier.GetBatchByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetBatchByID.Lock()
	mock.calls.GetBatchByID = append(mock.calls.GetBatchByID, callInfo)
	mock.lockGetBatchByID.Unlock()
	return mock.GetBatchByIDFunc(ctx, id)
}

// GetBatchByIDCalls gets all the calls that were made to GetBatchByID.
// Check the length with:
//
//	len(mockedQuerier.GetBatchByIDCalls())
func (mock *QuerierMock) GetBatchByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetBatchByID.RLock()
	calls = mock.calls.GetBatchByID
	mock.lockGetBatchByID.RUnlock()
	return calls
}

// GetBatchRowsByBatchID calls GetBatchRowsByBatchIDFunc.
func (mock *QuerierMock) GetBatchRowsByBatchID(ctx context.Context, batch uuid.UUID) ([]batchsqlc.Batchrow, error) {
	if mock.GetBatchRowsByBatchIDFunc == nil {
		panic("QuerierMock.GetBatchRowsByBatchIDFunc: method is nil but Querier.GetBatchRowsByBatchID was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockGetBatchRowsByBatchID.Lock()
	mock.calls.GetBatchRowsByBatchID = append(mock.calls.GetBatchRowsByBatchID, callInfo)
	mock.lockGetBatchRowsByBatchID.Unlock()
	return mock.GetBatchRowsByBatchIDFunc(ctx, batch)
}

// GetBatchRowsByBatchIDCalls gets all the calls that were made to GetBatchRowsByBatchID.
// Check the length with:
//
//	len(mockedQuerier.GetBatchRowsByBatchIDCalls())
func (mock *QuerierMock) GetBatchRowsByBatchIDCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockGetBatchRowsByBatchID.RLock()
	calls = mock.calls.GetBatchRowsByBatchID
	mock.lockGetBatchRowsByBatchID.RUnlock()
	return calls
}

// GetBatchRowsByBatchIDSorted calls GetBatchRowsByBatchIDSortedFunc.
func (mock *QuerierMock) GetBatchRowsByBatchIDSorted(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetBatchRowsByBatchIDSortedRow, error) {
	if mock.GetBatchRowsByBatchIDSortedFunc == nil {
		panic("QuerierMock.GetBatchRowsByBatchIDSortedFunc: method is nil but Querier.GetBatchRowsByBatchIDSorted was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockGetBatchRowsByBatchIDSorted.Lock()
	mock.calls.GetBatchRowsByBatchIDSorted = append(mock.calls.GetBatchRowsByBatchIDSorted, callInfo)
	mock.lockGetBatchRowsByBatchIDSorted.Unlock()
	return mock.GetBatchRowsByBatchIDSortedFunc(ctx, batch)
}

// GetBatchRowsByBatchIDSortedCalls gets all the calls that were made to GetBatchRowsByBatchIDSorted.
// Check the length with:
//
//	len(mockedQuerier.GetBatchRowsByBatchIDSortedCalls())
func (mock *QuerierMock) GetBatchRowsByBatchIDSortedCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockGetBatchRowsByBatchIDSorted.RLock()
	calls = mock.calls.GetBatchRowsByBatchIDSorted
	mock.lockGetBatchRowsByBatchIDSorted.RUnlock()
	return calls
}

// GetBatchRowsCount calls GetBatchRowsCountFunc.
func (mock *QuerierMock) GetBatchRowsCount(ctx context.Context, batch uuid.UUID) (int64, error) {
	if mock.GetBatchRowsCountFunc == nil {
		panic("QuerierMock.GetBatchRowsCountFunc: method is nil but Querier.GetBatchRowsCount was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockGetBatchRowsCount.Lock()
	mock.calls.GetBatchRowsCount = append(mock.calls.GetBatchRowsCount, callInfo)
	mock.lockGetBatchRowsCount.Unlock()
	return mock.GetBatchRowsCountFunc(ctx, batch)
}

// GetBatchRowsCountCalls gets all the calls that were made to GetBatchRowsCount.
// Check the length with:
//
//	len(mockedQuerier.GetBatchRowsCountCalls())
func (mock *QuerierMock) GetBatchRowsCountCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockGetBatchRowsCount.RLock()
	calls = mock.calls.GetBatchRowsCount
	mock.lockGetBatchRowsCount.RUnlock()
	return calls
}

// GetBatchStatus calls GetBatchStatusFunc.
func (mock *QuerierMock) GetBatchStatus(ctx context.Context, id uuid.UUID) (batchsqlc.StatusEnum, error) {
	if mock.GetBatchStatusFunc == nil {
		panic("QuerierMock.GetBatchStatusFunc: method is nil but Querier.GetBatchStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetBatchStatus.Lock()
	mock.calls.GetBatchStatus = append(mock.calls.GetBatchStatus, callInfo)
	mock.lockGetBatchStatus.Unlock()
	return mock.GetBatchStatusFunc(ctx, id)
}

// GetBatchStatusCalls gets all the calls that were made to GetBatchStatus.
// Check the length with:
//
//	len(mockedQuerier.GetBatchStatusCalls())
func (mock *QuerierMock) GetBatchStatusCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetBatchStatus.RLock()
	calls = mock.calls.GetBatchStatus
	mock.lockGetBatchStatus.RUnlock()
	return calls
}

// GetCompletedBatches calls GetCompletedBatchesFunc.
func (mock *QuerierMock) GetCompletedBatches(ctx context.Context) ([]uuid.UUID, error) {
	if mock.GetCompletedBatchesFunc == nil {
		panic("QuerierMock.GetCompletedBatchesFunc: method is nil but Querier.GetCompletedBatches was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetCompletedBatches.Lock()
	mock.calls.GetCompletedBatches = append(mock.calls.GetCompletedBatches, callInfo)
	mock.lockGetCompletedBatches.Unlock()
	return mock.GetCompletedBatchesFunc(ctx)
}

// GetCompletedBatchesCalls gets all the calls that were made to GetCompletedBatches.
// Check the length with:
//
//	len(mockedQuerier.GetCompletedBatchesCalls())
func (mock *QuerierMock) GetCompletedBatchesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetCompletedBatches.RLock()
	calls = mock.calls.GetCompletedBatches
	mock.lockGetCompletedBatches.RUnlock()
	return calls
}

// GetPendingBatchRows calls GetPendingBatchRowsFunc.
func (mock *QuerierMock) GetPendingBatchRows(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetPendingBatchRowsRow, error) {
	if mock.GetPendingBatchRowsFunc == nil {
		panic("QuerierMock.GetPendingBatchRowsFunc: method is nil but Querier.GetPendingBatchRows was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockGetPendingBatchRows.Lock()
	mock.calls.GetPendingBatchRows = append(mock.calls.GetPendingBatchRows, callInfo)
	mock.lockGetPendingBatchRows.Unlock()
	return mock.GetPendingBatchRowsFunc(ctx, batch)
}

// GetPendingBatchRowsCalls gets all the calls that were made to GetPendingBatchRows.
// Check the length with:
//
//	len(mockedQuerier.GetPendingBatchRowsCalls())
func (mock *QuerierMock) GetPendingBatchRowsCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockGetPendingBatchRows.RLock()
	calls = mock.calls.GetPendingBatchRows
	mock.lockGetPendingBatchRows.RUnlock()
	return calls
}

// GetProcessedBatchRowsByBatchIDSorted calls GetProcessedBatchRowsByBatchIDSortedFunc.
func (mock *QuerierMock) GetProcessedBatchRowsByBatchIDSorted(ctx context.Context, batch uuid.UUID) ([]batchsqlc.GetProcessedBatchRowsByBatchIDSortedRow, error) {
	if mock.GetProcessedBatchRowsByBatchIDSortedFunc == nil {
		panic("QuerierMock.GetProcessedBatchRowsByBatchIDSortedFunc: method is nil but Querier.GetProcessedBatchRowsByBatchIDSorted was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch uuid.UUID
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockGetProcessedBatchRowsByBatchIDSorted.Lock()
	mock.calls.GetProcessedBatchRowsByBatchIDSorted = append(mock.calls.GetProcessedBatchRowsByBatchIDSorted, callInfo)
	mock.lockGetProcessedBatchRowsByBatchIDSorted.Unlock()
	return mock.GetProcessedBatchRowsByBatchIDSortedFunc(ctx, batch)
}

// GetProcessedBatchRowsByBatchIDSortedCalls gets all the calls that were made to GetProcessedBatchRowsByBatchIDSorted.
// Check the length with:
//
//	len(mockedQuerier.GetProcessedBatchRowsByBatchIDSortedCalls())
func (mock *QuerierMock) GetProcessedBatchRowsByBatchIDSortedCalls() []struct {
	Ctx   context.Context
	Batch uuid.UUID
} {
	var calls []struct {
		Ctx   context.Context
		Batch uuid.UUID
	}
	mock.lockGetProcessedBatchRowsByBatchIDSorted.RLock()
	calls = mock.calls.GetProcessedBatchRowsByBatchIDSorted
	mock.lockGetProcessedBatchRowsByBatchIDSorted.RUnlock()
	return calls
}

// InsertIntoBatchRows calls InsertIntoBatchRowsFunc.
func (mock *QuerierMock) InsertIntoBatchRows(ctx context.Context, arg batchsqlc.InsertIntoBatchRowsParams) error {
	if mock.InsertIntoBatchRowsFunc == nil {
		panic("QuerierMock.InsertIntoBatchRowsFunc: method is nil but Querier.InsertIntoBatchRows was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.InsertIntoBatchRowsParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockInsertIntoBatchRows.Lock()
	mock.calls.InsertIntoBatchRows = append(mock.calls.InsertIntoBatchRows, callInfo)
	mock.lockInsertIntoBatchRows.Unlock()
	return mock.InsertIntoBatchRowsFunc(ctx, arg)
}

// InsertIntoBatchRowsCalls gets all the calls that were made to InsertIntoBatchRows.
// Check the length with:
//
//	len(mockedQuerier.InsertIntoBatchRowsCalls())
func (mock *QuerierMock) InsertIntoBatchRowsCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.InsertIntoBatchRowsParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.InsertIntoBatchRowsParams
	}
	mock.lockInsertIntoBatchRows.RLock()
	calls = mock.calls.InsertIntoBatchRows
	mock.lockInsertIntoBatchRows.RUnlock()
	return calls
}

// InsertIntoBatches calls InsertIntoBatchesFunc.
func (mock *QuerierMock) InsertIntoBatches(ctx context.Context, arg batchsqlc.InsertIntoBatchesParams) (uuid.UUID, error) {
	if mock.InsertIntoBatchesFunc == nil {
		panic("QuerierMock.InsertIntoBatchesFunc: method is nil but Querier.InsertIntoBatches was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.InsertIntoBatchesParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockInsertIntoBatches.Lock()
	mock.calls.InsertIntoBatches = append(mock.calls.InsertIntoBatches, callInfo)
	mock.lockInsertIntoBatches.Unlock()
	return mock.InsertIntoBatchesFunc(ctx, arg)
}

// InsertIntoBatchesCalls gets all the calls that were made to InsertIntoBatches.
// Check the length with:
//
//	len(mockedQuerier.InsertIntoBatchesCalls())
func (mock *QuerierMock) InsertIntoBatchesCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.InsertIntoBatchesParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.InsertIntoBatchesParams
	}
	mock.lockInsertIntoBatches.RLock()
	calls = mock.calls.InsertIntoBatches
	mock.lockInsertIntoBatches.RUnlock()
	return calls
}

// UpdateBatchCounters calls UpdateBatchCountersFunc.
func (mock *QuerierMock) UpdateBatchCounters(ctx context.Context, arg batchsqlc.UpdateBatchCountersParams) error {
	if mock.UpdateBatchCountersFunc == nil {
		panic("QuerierMock.UpdateBatchCountersFunc: method is nil but Querier.UpdateBatchCounters was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchCountersParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchCounters.Lock()
	mock.calls.UpdateBatchCounters = append(mock.calls.UpdateBatchCounters, callInfo)
	mock.lockUpdateBatchCounters.Unlock()
	return mock.UpdateBatchCountersFunc(ctx, arg)
}

// UpdateBatchCountersCalls gets all the calls that were made to UpdateBatchCounters.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchCountersCalls())
func (mock *QuerierMock) UpdateBatchCountersCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchCountersParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchCountersParams
	}
	mock.lockUpdateBatchCounters.RLock()
	calls = mock.calls.UpdateBatchCounters
	mock.lockUpdateBatchCounters.RUnlock()
	return calls
}

// UpdateBatchOutputFiles calls UpdateBatchOutputFilesFunc.
func (mock *QuerierMock) UpdateBatchOutputFiles(ctx context.Context, arg batchsqlc.UpdateBatchOutputFilesParams) error {
	if mock.UpdateBatchOutputFilesFunc == nil {
		panic("QuerierMock.UpdateBatchOutputFilesFunc: method is nil but Querier.UpdateBatchOutputFiles was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchOutputFilesParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchOutputFiles.Lock()
	mock.calls.UpdateBatchOutputFiles = append(mock.calls.UpdateBatchOutputFiles, callInfo)
	mock.lockUpdateBatchOutputFiles.Unlock()
	return mock.UpdateBatchOutputFilesFunc(ctx, arg)
}

// UpdateBatchOutputFilesCalls gets all the calls that were made to UpdateBatchOutputFiles.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchOutputFilesCalls())
func (mock *QuerierMock) UpdateBatchOutputFilesCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchOutputFilesParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchOutputFilesParams
	}
	mock.lockUpdateBatchOutputFiles.RLock()
	calls = mock.calls.UpdateBatchOutputFiles
	mock.lockUpdateBatchOutputFiles.RUnlock()
	return calls
}

// UpdateBatchRowsBatchJob calls UpdateBatchRowsBatchJobFunc.
func (mock *QuerierMock) UpdateBatchRowsBatchJob(ctx context.Context, arg batchsqlc.UpdateBatchRowsBatchJobParams) error {
	if mock.UpdateBatchRowsBatchJobFunc == nil {
		panic("QuerierMock.UpdateBatchRowsBatchJobFunc: method is nil but Querier.UpdateBatchRowsBatchJob was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsBatchJobParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchRowsBatchJob.Lock()
	mock.calls.UpdateBatchRowsBatchJob = append(mock.calls.UpdateBatchRowsBatchJob, callInfo)
	mock.lockUpdateBatchRowsBatchJob.Unlock()
	return mock.UpdateBatchRowsBatchJobFunc(ctx, arg)
}

// UpdateBatchRowsBatchJobCalls gets all the calls that were made to UpdateBatchRowsBatchJob.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchRowsBatchJobCalls())
func (mock *QuerierMock) UpdateBatchRowsBatchJobCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchRowsBatchJobParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsBatchJobParams
	}
	mock.lockUpdateBatchRowsBatchJob.RLock()
	calls = mock.calls.UpdateBatchRowsBatchJob
	mock.lockUpdateBatchRowsBatchJob.RUnlock()
	return calls
}

// UpdateBatchRowsSlowQuery calls UpdateBatchRowsSlowQueryFunc.
func (mock *QuerierMock) UpdateBatchRowsSlowQuery(ctx context.Context, arg batchsqlc.UpdateBatchRowsSlowQueryParams) error {
	if mock.UpdateBatchRowsSlowQueryFunc == nil {
		panic("QuerierMock.UpdateBatchRowsSlowQueryFunc: method is nil but Querier.UpdateBatchRowsSlowQuery was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsSlowQueryParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchRowsSlowQuery.Lock()
	mock.calls.UpdateBatchRowsSlowQuery = append(mock.calls.UpdateBatchRowsSlowQuery, callInfo)
	mock.lockUpdateBatchRowsSlowQuery.Unlock()
	return mock.UpdateBatchRowsSlowQueryFunc(ctx, arg)
}

// UpdateBatchRowsSlowQueryCalls gets all the calls that were made to UpdateBatchRowsSlowQuery.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchRowsSlowQueryCalls())
func (mock *QuerierMock) UpdateBatchRowsSlowQueryCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchRowsSlowQueryParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsSlowQueryParams
	}
	mock.lockUpdateBatchRowsSlowQuery.RLock()
	calls = mock.calls.UpdateBatchRowsSlowQuery
	mock.lockUpdateBatchRowsSlowQuery.RUnlock()
	return calls
}

// UpdateBatchRowsStatus calls UpdateBatchRowsStatusFunc.
func (mock *QuerierMock) UpdateBatchRowsStatus(ctx context.Context, arg batchsqlc.UpdateBatchRowsStatusParams) error {
	if mock.UpdateBatchRowsStatusFunc == nil {
		panic("QuerierMock.UpdateBatchRowsStatusFunc: method is nil but Querier.UpdateBatchRowsStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsStatusParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchRowsStatus.Lock()
	mock.calls.UpdateBatchRowsStatus = append(mock.calls.UpdateBatchRowsStatus, callInfo)
	mock.lockUpdateBatchRowsStatus.Unlock()
	return mock.UpdateBatchRowsStatusFunc(ctx, arg)
}

// UpdateBatchRowsStatusCalls gets all the calls that were made to UpdateBatchRowsStatus.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchRowsStatusCalls())
func (mock *QuerierMock) UpdateBatchRowsStatusCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchRowsStatusParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchRowsStatusParams
	}
	mock.lockUpdateBatchRowsStatus.RLock()
	calls = mock.calls.UpdateBatchRowsStatus
	mock.lockUpdateBatchRowsStatus.RUnlock()
	return calls
}

// UpdateBatchStatus calls UpdateBatchStatusFunc.
func (mock *QuerierMock) UpdateBatchStatus(ctx context.Context, arg batchsqlc.UpdateBatchStatusParams) error {
	if mock.UpdateBatchStatusFunc == nil {
		panic("QuerierMock.UpdateBatchStatusFunc: method is nil but Querier.UpdateBatchStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchStatusParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchStatus.Lock()
	mock.calls.UpdateBatchStatus = append(mock.calls.UpdateBatchStatus, callInfo)
	mock.lockUpdateBatchStatus.Unlock()
	return mock.UpdateBatchStatusFunc(ctx, arg)
}

// UpdateBatchStatusCalls gets all the calls that were made to UpdateBatchStatus.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchStatusCalls())
func (mock *QuerierMock) UpdateBatchStatusCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchStatusParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchStatusParams
	}
	mock.lockUpdateBatchStatus.RLock()
	calls = mock.calls.UpdateBatchStatus
	mock.lockUpdateBatchStatus.RUnlock()
	return calls
}

// UpdateBatchSummary calls UpdateBatchSummaryFunc.
func (mock *QuerierMock) UpdateBatchSummary(ctx context.Context, arg batchsqlc.UpdateBatchSummaryParams) error {
	if mock.UpdateBatchSummaryFunc == nil {
		panic("QuerierMock.UpdateBatchSummaryFunc: method is nil but Querier.UpdateBatchSummary was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchSummaryParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchSummary.Lock()
	mock.calls.UpdateBatchSummary = append(mock.calls.UpdateBatchSummary, callInfo)
	mock.lockUpdateBatchSummary.Unlock()
	return mock.UpdateBatchSummaryFunc(ctx, arg)
}

// UpdateBatchSummaryCalls gets all the calls that were made to UpdateBatchSummary.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchSummaryCalls())
func (mock *QuerierMock) UpdateBatchSummaryCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchSummaryParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchSummaryParams
	}
	mock.lockUpdateBatchSummary.RLock()
	calls = mock.calls.UpdateBatchSummary
	mock.lockUpdateBatchSummary.RUnlock()
	return calls
}

// UpdateBatchSummaryOnAbort calls UpdateBatchSummaryOnAbortFunc.
func (mock *QuerierMock) UpdateBatchSummaryOnAbort(ctx context.Context, arg batchsqlc.UpdateBatchSummaryOnAbortParams) error {
	if mock.UpdateBatchSummaryOnAbortFunc == nil {
		panic("QuerierMock.UpdateBatchSummaryOnAbortFunc: method is nil but Querier.UpdateBatchSummaryOnAbort was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchSummaryOnAbortParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockUpdateBatchSummaryOnAbort.Lock()
	mock.calls.UpdateBatchSummaryOnAbort = append(mock.calls.UpdateBatchSummaryOnAbort, callInfo)
	mock.lockUpdateBatchSummaryOnAbort.Unlock()
	return mock.UpdateBatchSummaryOnAbortFunc(ctx, arg)
}

// UpdateBatchSummaryOnAbortCalls gets all the calls that were made to UpdateBatchSummaryOnAbort.
// Check the length with:
//
//	len(mockedQuerier.UpdateBatchSummaryOnAbortCalls())
func (mock *QuerierMock) UpdateBatchSummaryOnAbortCalls() []struct {
	Ctx context.Context
	Arg batchsqlc.UpdateBatchSummaryOnAbortParams
} {
	var calls []struct {
		Ctx context.Context
		Arg batchsqlc.UpdateBatchSummaryOnAbortParams
	}
	mock.lockUpdateBatchSummaryOnAbort.RLock()
	calls = mock.calls.UpdateBatchSummaryOnAbort
	mock.lockUpdateBatchSummaryOnAbort.RUnlock()
	return calls
}
