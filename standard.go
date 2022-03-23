package imparse

import "context"

//
type Exec interface {
	Transport(ctx context.Context, id int64, key, from, target string, ch Channel, frameType FrameType, data []byte) error
	RevAck(ctx context.Context, id int64, keys []string, data []byte) error
}

type Filter func(ctx context.Context, frame Frame) error
type ActionFilter interface {
	FrameFilters(tp FrameType) []Filter
}

//implement
type StandardAnswer struct {
	cache        Cache
	exec         Exec
	trace        Trace
	frameFilters map[FrameType][]Filter
}

func NewStandardAnswer(cache Cache, exec Exec, trace Trace, filters map[FrameType][]Filter) *StandardAnswer {
	return &StandardAnswer{
		cache:        cache,
		exec:         exec,
		trace:        trace,
		frameFilters: filters,
	}
}

func (s *StandardAnswer) FrameFilters(tp FrameType) []Filter {
	if s.frameFilters == nil {
		s.frameFilters = make(map[FrameType][]Filter)
	}
	return s.frameFilters[tp]
}

func (s *StandardAnswer) Check(ctx context.Context, checker Checker, frame Frame) error {
	if s.trace != nil {
		finish, _ := s.trace.StartSpanFromContext(ctx, "Check")
		defer finish()
	}
	return checker.CheckFrame(frame)
}

func (s *StandardAnswer) Filter(ctx context.Context, frame Frame) (uint64, error) {
	if s.trace != nil {
		finish, _ := s.trace.StartSpanFromContext(ctx, "Filter")
		defer finish()
	}
	fs := s.FrameFilters(frame.Type())
	return frame.Filter(ctx, s.cache, fs...)
}

func (s *StandardAnswer) Transport(ctx context.Context, frame Frame) error {
	if s.trace != nil {
		finish, _ := s.trace.StartSpanFromContext(ctx, "Transport")
		defer finish()
	}
	return frame.Transport(ctx, s.exec)
}

func (s *StandardAnswer) Ack(ctx context.Context, frame Frame) (int64, error) {
	if s.trace != nil {
		finish, _ := s.trace.StartSpanFromContext(ctx, "Ack")
		defer finish()
	}
	return frame.Ack(ctx, s.exec)
}

//
type StandardStorage struct {
	db DB
}

func NewStandardStorage(db DB) *StandardStorage {
	return &StandardStorage{
		db: db,
	}
}

func (s *StandardStorage) SaveMsg(ctx context.Context, frame Frame) error {
	return s.db.SaveMsg(frame)
}
