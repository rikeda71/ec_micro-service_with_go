import React from "react";
import ReactDom from "react-dom";
import { Login } from "./components/login";

const App: React.FC = () => {
  return (
    <React.Fragment>
      <h1>EC Application with Micro Service</h1>
      <Login />
    </React.Fragment>
  );
};

ReactDom.render(<App />, document.getElementById("root") as HTMLElement);
