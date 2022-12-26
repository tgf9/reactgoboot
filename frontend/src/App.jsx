import React from "react";
import {
  createBrowserRouter,
  RouterProvider,
  Route,
} from "react-router-dom";

import Foo from "./Foo";

export default function App(props) {
  const r = createBrowserRouter([
    {
      path: "/",
      element: <Foo />,
    },
  ]);

  return <RouterProvider router={r} />;
};
