import React, { useState } from 'react';
import ReactDom from 'react-dom';
import { Login } from './components/login';
import { Products } from './components/products';

const App: React.FC = () => {
  const [isLogin, setIsLogin] = useState(false);
  return (
    <React.Fragment>
      <h1>EC Application with Micro Service</h1>
      <Login isLogin={isLogin} setIsLogin={setIsLogin} />
      <Products isLogin={isLogin} />
    </React.Fragment>
  );
};

ReactDom.render(<App />, document.getElementById('root') as HTMLElement);
