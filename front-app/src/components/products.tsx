import React, { useState } from 'react';
import styled from 'styled-components';
import Cookies from 'js-cookie';

export interface IProduct {
  product_id: number;
  product_name: string;
  product_image: string;
  product_price: number;
}

const ProductsStyle = styled.div`
  em {
    font-size: 1rem;
    color: #f00;
  }
`;

export const Products: React.FC<{ isLogin: boolean; products: Array<IProduct>; setProducts: Function }> = (props) => {
  const fetchCarts = () => {
    const method = 'GET';
    const headers = { Accept: 'application/json' };
    fetch('http://localhost:3001/products', { method, headers })
      .then(function (resp) {
        return resp.json();
      })
      .then(function (json) {
        if (json) {
          props.setProducts(json);
        }
      });
  };
  fetchCarts();

  return (
    <ProductsStyle>
      <h2>商品一覧</h2>
      <table width="600">
        <tr>
          <td>
            {props.products.map((p) => (
              <Product product={p} isLogin={props.isLogin} />
            ))}
          </td>
        </tr>
      </table>
    </ProductsStyle>
  );
};

const Product: React.FC<{ product: IProduct; isLogin: boolean }> = (props) => {
  const cartIn = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    const token = Cookies.get('token');
    if (token) {
      const method = 'POST';
      const headers = {
        Accept: 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
        Authorization: 'Bearer ' + token,
      };
      const body = `product_id=${props.product.product_id}`;
      fetch('http://localhost:3002/cart', { method, headers, body })
        .then(function (resp) {
          return resp.json();
        })
        .then(function (json) {
          if (json) {
            console.log(json);
          }
        });
    }
  };

  return (
    <React.Fragment>
      <div>
        <img src={props.product.product_image} alt={props.product.product_name} width="200" />
      </div>
      <div>ProductName:{props.product.product_name}</div>
      <div>Price:{props.product.product_price}</div>
      {props.isLogin && <button onClick={cartIn}>カートに入れる</button>}
    </React.Fragment>
  );
};
