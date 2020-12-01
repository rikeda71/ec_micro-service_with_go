import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import Cookies from 'js-cookie';
import { IProduct } from './products';

interface ICart {
  product_id: number;
  user_id: number;
}

interface IProps {
  products: Array<IProduct>;
  cartNum: number;
  setCartNum: Function;
}

const CartStyle = styled.div`
  em {
    font-size: 1rem;
    color: #f00;
  }
`;

export const Cart: React.FC<IProps> = (props) => {
  const [carts, setCarts] = useState<Array<IProduct>>([]);
  const [totalCost, setTotalCost] = useState<number>(0);
  const token = Cookies.get('token');

  const fetchCartItems = () => {
    const method = 'GET';
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
      Authorization: 'Bearer ' + token,
    };
    const callFetch = async () => {
      await fetch('http://localhost:3002/cart', { method, headers })
        .then(function (resp) {
          return resp.json();
        })
        .then(function (json) {
          if (json) {
            let cost = 0;
            console.log(json);
            const newCarts: Array<IProduct> = [];
            json.forEach((cart: ICart) => {
              props.products.forEach(function (product) {
                if (product.product_id === cart.product_id) {
                  newCarts.push(product);
                  // 合計金額
                  cost += product.product_price;
                  setTotalCost(cost);
                }
              });
            });
            setCarts([...newCarts]);
          }
        });
    };
    callFetch();
  };
  //------------------------- // 商品購入 //-------------------------
  const buy = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    const func = async () => {
      const method = 'POST';
      const headers = {
        Accept: 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
        Authorization: 'Bearer ' + token,
      };
      const body = JSON.stringify({ order_details: carts });
      await fetch('http://localhost:3003/order', { method, headers, body })
        .then(function (resp) {
          return resp.json();
        })
        .then(function (json) {
          if (json) {
            deleteCartItems();
            props.setCartNum(0);
          }
        });
    };
    func();
  };
  const deleteCartItems = () => {
    const method = 'DELETE';
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
          setCarts([]);
        }
      });
  };
  // クッキーがある場合は初期処理でカートアイテム取得
  if (token) {
    useEffect(() => {
      fetchCartItems();
    }, [props.products, props.cartNum]);
  }

  return (
    <CartStyle>
      <div>
        <h2>買い物かご</h2>
        {carts.length > 0 ? (
          <table style={{ border: '1px solid black;' }}>
            <tr>
              <th>商品名</th>
              <th>価格</th>
            </tr>
            {carts.map((c) => {
              return (
                <tr>
                  <td>{c.product_name}</td>
                  <td>{c.product_price}</td>
                </tr>
              );
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
        ) : (
          <div> 買い物カゴに商品はありません </div>
        )}
      </div>
    </CartStyle>
  );
};
