import React, { useState } from 'react';
import styled from 'styled-components';
import Cookies from 'js-cookie';
import { IProduct } from './products';

interface ICart {
  product_id: number;
  user_id: number;
}

const CartStyle = styled.div`
  em {
    font-size: 1rem;
    color: #f00;
  }
`;

export const Cart: React.FC<{ products: Array<IProduct> }> = (props) => {
  const [carts, setCarts] = useState<Array<IProduct>>([]);
  const [totalCost, setTotalCost] = useState(0);
  const token = Cookies.get('token');

  const fetchCartItems = () => {
    const method = 'GET';
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
      Authorization: 'Bearer ' + token,
    };
    fetch('http://localhost:3002/cart', { method, headers })
      .then(function (resp) {
        return resp.json();
      })
      .then(function (json) {
        if (json) {
          console.log(json);
          json.forEach((cart: ICart) => {
            props.products.forEach(function (product) {
              if (product.product_id == cart.product_id) {
                setCarts(carts.concat([product]));
              }
            });
          });
          // 合計金額
          carts.forEach((c) => {
            setTotalCost(totalCost + c.product_price);
          });
        }
      });
  };
  //------------------------- // 商品購入 //-------------------------
  const buy = () => {
    const method = 'POST';
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
      Authorization: 'Bearer ' + token,
    };
    const body = JSON.stringify({ order_details: carts });
    fetch('http://localhost:3003/order', { method, headers, body })
      .then(function (resp) {
        return resp.json();
      })
      .then(function (json) {
        if (json) {
          deleteCartItems();
        }
      });
  }; //------------------------- // カートアイテム全削除 //-------------------------
  const deleteCartItems = () => {
    const method = 'DELETE';
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
      Authorization: 'Bearer ' + token,
    };
    fetch('http://localhost:3002/carts', { method, headers })
      .then(function (resp) {
        return resp.json();
      })
      .then(function (json) {
        if (json) {
          fetchCartItems();
        }
      });
  }; // クッキーある場合は初期処理でカートアイテム取得
  if (token) {
    fetchCartItems();
  }

  return (
    <CartStyle>
      <div>
        <h2>買い物かご</h2>
        {carts.length > 0 && (
          <table>
            <tr>
              <th>商品名</th>
              <th>価格</th>
            </tr>
            {carts.map((c) => {
              <tr>
                <React.Fragment>
                  <td>{c.product_name}</td>
                  <td>{c.product_price}</td>
                </React.Fragment>
              </tr>;
            })}
            <tr>
              <td colSpan={2}>
                <div className="text-center">
                  <b>合計金額: {totalCost}円</b>
                </div>
              </td>
            </tr>
            <tr>
              <td colSpan={2}>
                <button onClick={buy}>購入</button>
              </td>
            </tr>
          </table>
        )}
        {carts.length == 0 && <div> 買い物カゴに商品はありません </div>}
      </div>
    </CartStyle>
  );
};
