import React, { useEffect } from 'react';
import styled from 'styled-components';
import Cookies from 'js-cookie';

export interface IProduct {
  product_id: number;
  product_name: string;
  product_image: string;
  product_price: number;
}

interface IProps {
  isLogin: boolean;
  products: Array<IProduct>;
  setProducts: Function;
  cartNum: number;
  setCartNum: Function;
}

const ProductsStyle = styled.div`
  em {
    font-size: 1rem;
    color: #f00;
  }
`;

export const Products: React.FC<IProps> = (props) => {
  useEffect(() => {
    const fetchProduct = async () => {
      const method = 'GET';
      const headers = { Accept: 'application/json' };
      await fetch('http://localhost:3001/products', { method, headers })
        .then(function (resp) {
          return resp.json();
        })
        .then(function (json) {
          if (json) {
            console.log(json);
            props.setProducts(json);
          }
        });
    };
    fetchProduct();
  }, []);

  return (
    <ProductsStyle>
      <h2>商品一覧</h2>
      <table width="600">
        <tr>
          <td>
            {props.products.map((p) => (
              <Product product={p} isLogin={props.isLogin} cartNum={props.cartNum} setCartNum={props.setCartNum} />
            ))}
          </td>
        </tr>
      </table>
    </ProductsStyle>
  );
};

const Product: React.FC<{ product: IProduct; isLogin: boolean; cartNum: number; setCartNum: Function }> = (props) => {
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
      const callFetch = async () => {
        await fetch('http://localhost:3002/cart', { method, headers, body })
          .then(function (resp) {
            return resp.json();
          })
          .then(function (json) {
            if (json) {
              console.log(json);
              props.setCartNum(props.cartNum + 1);
            }
          });
      };
      callFetch();
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
