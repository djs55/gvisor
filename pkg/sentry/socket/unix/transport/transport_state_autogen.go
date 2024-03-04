// automatically generated by stateify.

package transport

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (e *connectionedEndpoint) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.connectionedEndpoint"
}

func (e *connectionedEndpoint) StateFields() []string {
	return []string{
		"baseEndpoint",
		"id",
		"idGenerator",
		"stype",
		"acceptedChan",
		"boundSocketFD",
	}
}

// +checklocksignore
func (e *connectionedEndpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	var acceptedChanValue []*connectionedEndpoint
	acceptedChanValue = e.saveAcceptedChan()
	stateSinkObject.SaveValue(4, acceptedChanValue)
	stateSinkObject.Save(0, &e.baseEndpoint)
	stateSinkObject.Save(1, &e.id)
	stateSinkObject.Save(2, &e.idGenerator)
	stateSinkObject.Save(3, &e.stype)
	stateSinkObject.Save(5, &e.boundSocketFD)
}

// +checklocksignore
func (e *connectionedEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.baseEndpoint)
	stateSourceObject.Load(1, &e.id)
	stateSourceObject.Load(2, &e.idGenerator)
	stateSourceObject.Load(3, &e.stype)
	stateSourceObject.Load(5, &e.boundSocketFD)
	stateSourceObject.LoadValue(4, new([]*connectionedEndpoint), func(y any) { e.loadAcceptedChan(y.([]*connectionedEndpoint)) })
	stateSourceObject.AfterLoad(func() { e.afterLoad(ctx) })
}

func (e *connectionlessEndpoint) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.connectionlessEndpoint"
}

func (e *connectionlessEndpoint) StateFields() []string {
	return []string{
		"baseEndpoint",
	}
}

func (e *connectionlessEndpoint) beforeSave() {}

// +checklocksignore
func (e *connectionlessEndpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.baseEndpoint)
}

// +checklocksignore
func (e *connectionlessEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.baseEndpoint)
	stateSourceObject.AfterLoad(func() { e.afterLoad(ctx) })
}

func (c *HostConnectedEndpoint) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.HostConnectedEndpoint"
}

func (c *HostConnectedEndpoint) StateFields() []string {
	return []string{
		"HostConnectedEndpointRefs",
		"fd",
		"addr",
		"stype",
		"rdShutdown",
		"wrShutdown",
	}
}

func (c *HostConnectedEndpoint) beforeSave() {}

// +checklocksignore
func (c *HostConnectedEndpoint) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.HostConnectedEndpointRefs)
	stateSinkObject.Save(1, &c.fd)
	stateSinkObject.Save(2, &c.addr)
	stateSinkObject.Save(3, &c.stype)
	stateSinkObject.Save(4, &c.rdShutdown)
	stateSinkObject.Save(5, &c.wrShutdown)
}

// +checklocksignore
func (c *HostConnectedEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.HostConnectedEndpointRefs)
	stateSourceObject.Load(1, &c.fd)
	stateSourceObject.Load(2, &c.addr)
	stateSourceObject.Load(3, &c.stype)
	stateSourceObject.Load(4, &c.rdShutdown)
	stateSourceObject.Load(5, &c.wrShutdown)
	stateSourceObject.AfterLoad(func() { c.afterLoad(ctx) })
}

func (r *HostConnectedEndpointRefs) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.HostConnectedEndpointRefs"
}

func (r *HostConnectedEndpointRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *HostConnectedEndpointRefs) beforeSave() {}

// +checklocksignore
func (r *HostConnectedEndpointRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *HostConnectedEndpointRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (q *queue) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.queue"
}

func (q *queue) StateFields() []string {
	return []string{
		"queueRefs",
		"ReaderQueue",
		"WriterQueue",
		"closed",
		"unread",
		"used",
		"limit",
		"dataList",
	}
}

func (q *queue) beforeSave() {}

// +checklocksignore
func (q *queue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.queueRefs)
	stateSinkObject.Save(1, &q.ReaderQueue)
	stateSinkObject.Save(2, &q.WriterQueue)
	stateSinkObject.Save(3, &q.closed)
	stateSinkObject.Save(4, &q.unread)
	stateSinkObject.Save(5, &q.used)
	stateSinkObject.Save(6, &q.limit)
	stateSinkObject.Save(7, &q.dataList)
}

func (q *queue) afterLoad(context.Context) {}

// +checklocksignore
func (q *queue) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.queueRefs)
	stateSourceObject.Load(1, &q.ReaderQueue)
	stateSourceObject.Load(2, &q.WriterQueue)
	stateSourceObject.Load(3, &q.closed)
	stateSourceObject.Load(4, &q.unread)
	stateSourceObject.Load(5, &q.used)
	stateSourceObject.Load(6, &q.limit)
	stateSourceObject.Load(7, &q.dataList)
}

func (r *queueRefs) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.queueRefs"
}

func (r *queueRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *queueRefs) beforeSave() {}

// +checklocksignore
func (r *queueRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *queueRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (l *messageList) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.messageList"
}

func (l *messageList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *messageList) beforeSave() {}

// +checklocksignore
func (l *messageList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *messageList) afterLoad(context.Context) {}

// +checklocksignore
func (l *messageList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *messageEntry) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.messageEntry"
}

func (e *messageEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *messageEntry) beforeSave() {}

// +checklocksignore
func (e *messageEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *messageEntry) afterLoad(context.Context) {}

// +checklocksignore
func (e *messageEntry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (c *ControlMessages) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.ControlMessages"
}

func (c *ControlMessages) StateFields() []string {
	return []string{
		"Rights",
		"Credentials",
	}
}

func (c *ControlMessages) beforeSave() {}

// +checklocksignore
func (c *ControlMessages) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.Rights)
	stateSinkObject.Save(1, &c.Credentials)
}

func (c *ControlMessages) afterLoad(context.Context) {}

// +checklocksignore
func (c *ControlMessages) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.Rights)
	stateSourceObject.Load(1, &c.Credentials)
}

func (m *message) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.message"
}

func (m *message) StateFields() []string {
	return []string{
		"messageEntry",
		"Data",
		"Control",
		"Address",
	}
}

func (m *message) beforeSave() {}

// +checklocksignore
func (m *message) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.messageEntry)
	stateSinkObject.Save(1, &m.Data)
	stateSinkObject.Save(2, &m.Control)
	stateSinkObject.Save(3, &m.Address)
}

func (m *message) afterLoad(context.Context) {}

// +checklocksignore
func (m *message) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.messageEntry)
	stateSourceObject.Load(1, &m.Data)
	stateSourceObject.Load(2, &m.Control)
	stateSourceObject.Load(3, &m.Address)
}

func (a *Address) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.Address"
}

func (a *Address) StateFields() []string {
	return []string{
		"Addr",
	}
}

func (a *Address) beforeSave() {}

// +checklocksignore
func (a *Address) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Addr)
}

func (a *Address) afterLoad(context.Context) {}

// +checklocksignore
func (a *Address) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Addr)
}

func (q *queueReceiver) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.queueReceiver"
}

func (q *queueReceiver) StateFields() []string {
	return []string{
		"readQueue",
	}
}

func (q *queueReceiver) beforeSave() {}

// +checklocksignore
func (q *queueReceiver) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.readQueue)
}

func (q *queueReceiver) afterLoad(context.Context) {}

// +checklocksignore
func (q *queueReceiver) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.readQueue)
}

func (q *streamQueueReceiver) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.streamQueueReceiver"
}

func (q *streamQueueReceiver) StateFields() []string {
	return []string{
		"queueReceiver",
		"buffer",
		"control",
		"addr",
	}
}

func (q *streamQueueReceiver) beforeSave() {}

// +checklocksignore
func (q *streamQueueReceiver) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.queueReceiver)
	stateSinkObject.Save(1, &q.buffer)
	stateSinkObject.Save(2, &q.control)
	stateSinkObject.Save(3, &q.addr)
}

func (q *streamQueueReceiver) afterLoad(context.Context) {}

// +checklocksignore
func (q *streamQueueReceiver) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.queueReceiver)
	stateSourceObject.Load(1, &q.buffer)
	stateSourceObject.Load(2, &q.control)
	stateSourceObject.Load(3, &q.addr)
}

func (e *connectedEndpoint) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.connectedEndpoint"
}

func (e *connectedEndpoint) StateFields() []string {
	return []string{
		"endpoint",
		"writeQueue",
	}
}

func (e *connectedEndpoint) beforeSave() {}

// +checklocksignore
func (e *connectedEndpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.endpoint)
	stateSinkObject.Save(1, &e.writeQueue)
}

func (e *connectedEndpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *connectedEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.endpoint)
	stateSourceObject.Load(1, &e.writeQueue)
}

func (e *baseEndpoint) StateTypeName() string {
	return "pkg/sentry/socket/unix/transport.baseEndpoint"
}

func (e *baseEndpoint) StateFields() []string {
	return []string{
		"Queue",
		"DefaultSocketOptionsHandler",
		"receiver",
		"connected",
		"path",
		"ops",
	}
}

func (e *baseEndpoint) beforeSave() {}

// +checklocksignore
func (e *baseEndpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Queue)
	stateSinkObject.Save(1, &e.DefaultSocketOptionsHandler)
	stateSinkObject.Save(2, &e.receiver)
	stateSinkObject.Save(3, &e.connected)
	stateSinkObject.Save(4, &e.path)
	stateSinkObject.Save(5, &e.ops)
}

func (e *baseEndpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *baseEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Queue)
	stateSourceObject.Load(1, &e.DefaultSocketOptionsHandler)
	stateSourceObject.Load(2, &e.receiver)
	stateSourceObject.Load(3, &e.connected)
	stateSourceObject.Load(4, &e.path)
	stateSourceObject.Load(5, &e.ops)
}

func init() {
	state.Register((*connectionedEndpoint)(nil))
	state.Register((*connectionlessEndpoint)(nil))
	state.Register((*HostConnectedEndpoint)(nil))
	state.Register((*HostConnectedEndpointRefs)(nil))
	state.Register((*queue)(nil))
	state.Register((*queueRefs)(nil))
	state.Register((*messageList)(nil))
	state.Register((*messageEntry)(nil))
	state.Register((*ControlMessages)(nil))
	state.Register((*message)(nil))
	state.Register((*Address)(nil))
	state.Register((*queueReceiver)(nil))
	state.Register((*streamQueueReceiver)(nil))
	state.Register((*connectedEndpoint)(nil))
	state.Register((*baseEndpoint)(nil))
}
