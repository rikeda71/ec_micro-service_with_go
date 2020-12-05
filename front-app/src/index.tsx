import React, { useState } from 'react';
import ReactDom from 'react-dom';
import { Cart } from './components/cart';
import { Login } from './components/login';
import { Order } from './components/order';
import { Products } from './components/products';

const App: React.FC = () => {
  const [isLogin, setIsLogin] = useState(false);
  const [products, setProducts] = useState<Array<IProduct>>([]);
  const [cartNum, setCartNum] = useState<number>(0);
  return (
    <React.Fragment>
      <h1>EC Application with Micro Service</h1>
      <Login isLogin={isLogin} setIsLogin={setIsLogin} />
      <Products
        isLogin={isLogin}
        products={products}
        setProducts={setProducts}
        cartNum={cartNum}
        setCartNum={setCartNum}
      />
      {isLogin && <Cart products={products} cartNum={cartNum} setCartNum={setCartNum} />}
      {isLogin && <Order cartNum={cartNum} />}
    </React.Fragment>
  );
};

ReactDom.render(<App />, document.getElementById('root') as HTMLElement);
