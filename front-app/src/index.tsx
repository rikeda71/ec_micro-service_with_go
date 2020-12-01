import React, { useState } from 'react';
import ReactDom from 'react-dom';
import { Cart } from './components/cart';
import { Login } from './components/login';
import { IProduct, Products } from './components/products';

const App: React.FC = () => {
  const [isLogin, setIsLogin] = useState(false);
  const [products, setProducts] = useState<Array<IProduct>>([]);
  return (
    <React.Fragment>
      <h1>EC Application with Micro Service</h1>
      <Login isLogin={isLogin} setIsLogin={setIsLogin} />
      <Products isLogin={isLogin} products={products} setProducts={setProducts} />
      {isLogin && <Cart products={products} />}
    </React.Fragment>
  );
};

ReactDom.render(<App />, document.getElementById('root') as HTMLElement);
