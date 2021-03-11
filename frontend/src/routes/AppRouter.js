import React, { lazy, Suspense } from "react";
import { Switch, Route } from "react-router-dom";

const routes = [
  // Here's the route we defined
  {
    // The path that when the URL on the browser is match, then render this one
    path: "/",
    // The name of the route we define, this is optional
    name: "Homepage",
    // If the exact is false, then subpath will still let this component mounted on the views
    exact: true,
    // Children is the component will be mounted to the view when the url is matched
    // We use lazy function to load the component once the URL is matched, not eagerly load it.
    children: lazy(() => import("../views/Login")),
  },
  
  {
    path: "/login",
    name: "Counter",
    children: lazy(() => import("../views/Login")),
  },
];

const AppRouter = () => {
  // AppRouter is the way SPA like React render the specific component based on the URL
  // For more information, visit: https://reactrouter.com/web/guides/quick-start
  // Lazy load with React: https://reactjs.org/docs/code-splitting.html
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Switch>
        {routes.map((route) => (
          <Route key={route.path} {...route} children={<route.children />} />
        ))}
      </Switch>
    </Suspense>
  );
};

export default AppRouter;
