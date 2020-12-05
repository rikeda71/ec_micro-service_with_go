import Cookies from 'js-cookie';
import React, { useEffect, useState } from 'react';

interface IOrder {
  order_date: string;
  order_item_num: number;
  order_total_cost: number;
}

interface IOrderJson {
  CreatedAt: string;
  DeletedAt?: string;
  UpdatedAt: string;
  ID: number;
  user_id: number;
  OrderDetails: Array<IOrderDetail>;
}

export interface IOrderDetail {
  order_id?: number;
  product_id: number;
  product_price: number;
}

interface IProps {
  cartNum: number;
}

export const Order: React.FC<IProps> = (props) => {
  const [orders, setOrders] = useState<Array<IOrder>>([]);

  const token = Cookies.get('token');
  //-------------------------
  // 注文リスト取得
  //-------------------------
  const fetchOrders = () => {
    const method = 'GET';
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
      Authorization: 'Bearer ' + token,
    };
    const callFetch = async () => {
      await fetch('http://localhost:3003/orders', { method, headers })
        .then(function (resp) {
          return resp.json();
        })
        .then(function (json: Array<IOrderJson>) {
          if (json) {
            console.log(json);
            const newOrders: Array<IOrder> = [];
            json.forEach(function (order: IOrderJson) {
              let order_total_cost = 0;
              order.OrderDetails.forEach(function (order_detail) {
                // 合計金額計算
                order_total_cost += order_detail.product_price;
              });
              newOrders.push({
                order_date: order.CreatedAt,
                order_item_num: order.OrderDetails.length,
                order_total_cost: order_total_cost,
              });
            });
            setOrders(newOrders);
          }
        });
    };
    callFetch();
  };
  // クッキーある場合は初期処理で注文取得
  if (token) {
    useEffect(() => {
      fetchOrders();
    }, [token, props.cartNum]);
  }

  return (
    <div>
      <h2>注文一覧</h2>
      {orders.length != 0 ? (
        <table>
          <tr>
            <th>注文日時</th>
            <th>商品数</th>
            <th>金額</th>
          </tr>
          {orders.map((order) => {
            return (
              <tr>
                <td>{order.order_date}</td>
                <td>{order.order_item_num}</td>
                <td>{order.order_total_cost}</td>
              </tr>
            );
          })}
        </table>
      ) : (
        <div>注文はありません</div>
      )}
    </div>
  );
};
