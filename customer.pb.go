// Code generated by protoc-gen-go.
// source: customer.proto
// DO NOT EDIT!

/*
Package customerslist is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	Client
	Customers
	Customer
*/
package customerslist

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// client structure
type Client struct {
	Idc   int32  `protobuf:"varint,1,opt,name=idc" json:"idc,omitempty"`
	Namec string `protobuf:"bytes,2,opt,name=namec" json:"namec,omitempty"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}

// customer structure
type Customers struct {
	// int32 id=1;
	// string name=2;
	// string client=3;
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Customers) Reset()         { *m = Customers{} }
func (m *Customers) String() string { return proto.CompactTextString(m) }
func (*Customers) ProtoMessage()    {}

type Customer struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}

func init() {
}

// Client API for CustomersList service

type CustomersListClient interface {
	// add customers
	AddCustomers(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customers, error)
}

type customersListClient struct {
	cc *grpc.ClientConn
}

func NewCustomersListClient(cc *grpc.ClientConn) CustomersListClient {
	return &customersListClient{cc}
}

func (c *customersListClient) AddCustomers(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customers, error) {
	out := new(Customers)
	err := grpc.Invoke(ctx, "/customerslist.CustomersList/AddCustomers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomersList service

type CustomersListServer interface {
	// add customers
	AddCustomers(context.Context, *Customer) (*Customers, error)
}

func RegisterCustomersListServer(s *grpc.Server, srv CustomersListServer) {
	s.RegisterService(&_CustomersList_serviceDesc, srv)
}

func _CustomersList_AddCustomers_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Customer)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CustomersListServer).AddCustomers(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CustomersList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customerslist.CustomersList",
	HandlerType: (*CustomersListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCustomers",
			Handler:    _CustomersList_AddCustomers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
