// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/order.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "CrocsClub/pkg/domain"
	models "CrocsClub/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// AddOrderProducts mocks base method.
func (m *MockOrderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrderProducts", order_id, cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrderProducts indicates an expected call of AddOrderProducts.
func (mr *MockOrderRepositoryMockRecorder) AddOrderProducts(order_id, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrderProducts", reflect.TypeOf((*MockOrderRepository)(nil).AddOrderProducts), order_id, cart)
}

// ApproveOrder mocks base method.
func (m *MockOrderRepository) ApproveOrder(orderId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveOrder", orderId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveOrder indicates an expected call of ApproveOrder.
func (mr *MockOrderRepositoryMockRecorder) ApproveOrder(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveOrder", reflect.TypeOf((*MockOrderRepository)(nil).ApproveOrder), orderId)
}

// CancelOrder mocks base method.
func (m *MockOrderRepository) CancelOrder(orderId, userId, cartAmt int, paymentStatus string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", orderId, userId, cartAmt, paymentStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockOrderRepositoryMockRecorder) CancelOrder(orderId, userId, cartAmt, paymentStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockOrderRepository)(nil).CancelOrder), orderId, userId, cartAmt, paymentStatus)
}

// ChangeOrderStatus mocks base method.
func (m *MockOrderRepository) ChangeOrderStatus(orderID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", orderID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockOrderRepositoryMockRecorder) ChangeOrderStatus(orderID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockOrderRepository)(nil).ChangeOrderStatus), orderID, status)
}

// CheckOrderStatusByID mocks base method.
func (m *MockOrderRepository) CheckOrderStatusByID(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOrderStatusByID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOrderStatusByID indicates an expected call of CheckOrderStatusByID.
func (mr *MockOrderRepositoryMockRecorder) CheckOrderStatusByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOrderStatusByID", reflect.TypeOf((*MockOrderRepository)(nil).CheckOrderStatusByID), id)
}

// CheckOrdersStatusByID mocks base method.
func (m *MockOrderRepository) CheckOrdersStatusByID(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOrdersStatusByID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOrdersStatusByID indicates an expected call of CheckOrdersStatusByID.
func (mr *MockOrderRepositoryMockRecorder) CheckOrdersStatusByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOrdersStatusByID", reflect.TypeOf((*MockOrderRepository)(nil).CheckOrdersStatusByID), id)
}

// DebitWallet mocks base method.
func (m *MockOrderRepository) DebitWallet(userID int, Amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebitWallet", userID, Amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DebitWallet indicates an expected call of DebitWallet.
func (mr *MockOrderRepositoryMockRecorder) DebitWallet(userID, Amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebitWallet", reflect.TypeOf((*MockOrderRepository)(nil).DebitWallet), userID, Amount)
}

// GetAllOrders mocks base method.
func (m *MockOrderRepository) GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders", userId, page, pageSize)
	ret0, _ := ret[0].([]models.OrderDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockOrderRepositoryMockRecorder) GetAllOrders(userId, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockOrderRepository)(nil).GetAllOrders), userId, page, pageSize)
}

// GetDetailedOrderThroughId mocks base method.
func (m *MockOrderRepository) GetDetailedOrderThroughId(orderId int) (models.ItemOrderDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailedOrderThroughId", orderId)
	ret0, _ := ret[0].(models.ItemOrderDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailedOrderThroughId indicates an expected call of GetDetailedOrderThroughId.
func (mr *MockOrderRepositoryMockRecorder) GetDetailedOrderThroughId(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailedOrderThroughId", reflect.TypeOf((*MockOrderRepository)(nil).GetDetailedOrderThroughId), orderId)
}

// GetItemsByOrderId mocks base method.
func (m *MockOrderRepository) GetItemsByOrderId(orderId int) ([]models.ItemDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemsByOrderId", orderId)
	ret0, _ := ret[0].([]models.ItemDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemsByOrderId indicates an expected call of GetItemsByOrderId.
func (mr *MockOrderRepositoryMockRecorder) GetItemsByOrderId(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemsByOrderId", reflect.TypeOf((*MockOrderRepository)(nil).GetItemsByOrderId), orderId)
}

// GetOrderDetailsBrief mocks base method.
func (m *MockOrderRepository) GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetailsBrief", page)
	ret0, _ := ret[0].([]models.CombinedOrderDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetailsBrief indicates an expected call of GetOrderDetailsBrief.
func (mr *MockOrderRepositoryMockRecorder) GetOrderDetailsBrief(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetailsBrief", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderDetailsBrief), page)
}

// GetOrderDetailsByOrderId mocks base method.
func (m *MockOrderRepository) GetOrderDetailsByOrderId(orderID int) (models.CombinedOrderDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetailsByOrderId", orderID)
	ret0, _ := ret[0].(models.CombinedOrderDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetailsByOrderId indicates an expected call of GetOrderDetailsByOrderId.
func (mr *MockOrderRepositoryMockRecorder) GetOrderDetailsByOrderId(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetailsByOrderId", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderDetailsByOrderId), orderID)
}

// GetOrders mocks base method.
func (m *MockOrderRepository) GetOrders(orderId int) (domain.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", orderId)
	ret0, _ := ret[0].(domain.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepositoryMockRecorder) GetOrders(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepository)(nil).GetOrders), orderId)
}

// GetShipmentStatus mocks base method.
func (m *MockOrderRepository) GetShipmentStatus(orderId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShipmentStatus", orderId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShipmentStatus indicates an expected call of GetShipmentStatus.
func (mr *MockOrderRepositoryMockRecorder) GetShipmentStatus(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShipmentStatus", reflect.TypeOf((*MockOrderRepository)(nil).GetShipmentStatus), orderId)
}

// GetShipmentsStatus mocks base method.
func (m *MockOrderRepository) GetShipmentsStatus(orderID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShipmentsStatus", orderID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShipmentsStatus indicates an expected call of GetShipmentsStatus.
func (mr *MockOrderRepositoryMockRecorder) GetShipmentsStatus(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShipmentsStatus", reflect.TypeOf((*MockOrderRepository)(nil).GetShipmentsStatus), orderID)
}

// OrderItems mocks base method.
func (m *MockOrderRepository) OrderItems(userid, addressid, paymentid int, total float64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderItems", userid, addressid, paymentid, total)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderItems indicates an expected call of OrderItems.
func (mr *MockOrderRepositoryMockRecorder) OrderItems(userid, addressid, paymentid, total interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderItems", reflect.TypeOf((*MockOrderRepository)(nil).OrderItems), userid, addressid, paymentid, total)
}

// PaymentAlreadyPaid mocks base method.
func (m *MockOrderRepository) PaymentAlreadyPaid(orderID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentAlreadyPaid", orderID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaymentAlreadyPaid indicates an expected call of PaymentAlreadyPaid.
func (mr *MockOrderRepositoryMockRecorder) PaymentAlreadyPaid(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentAlreadyPaid", reflect.TypeOf((*MockOrderRepository)(nil).PaymentAlreadyPaid), orderID)
}

// PaymentMethodID mocks base method.
func (m *MockOrderRepository) PaymentMethodID(orderID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentMethodID", orderID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaymentMethodID indicates an expected call of PaymentMethodID.
func (mr *MockOrderRepositoryMockRecorder) PaymentMethodID(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentMethodID", reflect.TypeOf((*MockOrderRepository)(nil).PaymentMethodID), orderID)
}

// ReduceInventoryQuantity mocks base method.
func (m *MockOrderRepository) ReduceInventoryQuantity(productName string, quantity int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceInventoryQuantity", productName, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReduceInventoryQuantity indicates an expected call of ReduceInventoryQuantity.
func (mr *MockOrderRepositoryMockRecorder) ReduceInventoryQuantity(productName, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceInventoryQuantity", reflect.TypeOf((*MockOrderRepository)(nil).ReduceInventoryQuantity), productName, quantity)
}

// ReturnOrder mocks base method.
func (m *MockOrderRepository) ReturnOrder(returnOrderResp models.ReturnOrderResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnOrder", returnOrderResp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReturnOrder indicates an expected call of ReturnOrder.
func (mr *MockOrderRepositoryMockRecorder) ReturnOrder(returnOrderResp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnOrder", reflect.TypeOf((*MockOrderRepository)(nil).ReturnOrder), returnOrderResp)
}