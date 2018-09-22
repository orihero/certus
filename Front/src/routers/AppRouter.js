import React from "react";
import { Route, Switch, withRouter } from "react-router-dom";

import UsersPage from "../components/user/UsersPage";
import NotFoundPage from "../components/notFound/NotFoundPage";
import LoginPage from "../components/login/LoginPage";
import SeamstressPage from "../components/seamstress/SeamstressPage";
import SeamstressRoute from "../routes/SeamstressRoute";
import ScanPage from "../components/seamstress/ScanPage";
import EquipmentsPage from "../components/equipment/EquipmentsPage";
import TimerPage from "../components/seamstress/TimerPage";
import AdminRoute from "../routes/AdminRoute";

import "../assets/css/material-dashboard-react.css";
import "../assets/css/css.css";
import OperationsPage from "../components/operation/OperationsPage";

const AppRouter = ({ location }) => (
  <div className="wrapper">
    <Switch>
      <Route location={location} path="/login" component={LoginPage} />
      <SeamstressRoute path="/seamstress" component={SeamstressPage} />
      <Route path="/scan" component={ScanPage} />
      <Route location={location} path="/timer" component={TimerPage} />
      <Route
        location={location}
        path="/equipments"
        component={EquipmentsPage}
      />
      <Route location={location} path="/users" component={UsersPage} />
      <AdminRoute
        location={location}
        path="/operations"
        component={OperationsPage}
      />
      <Route location={location} component={NotFoundPage} />
    </Switch>
  </div>
);
export default withRouter(AppRouter);