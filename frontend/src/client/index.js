import React from "react";
import ReactDOM from "react-dom";
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import {TopPage} from "./component/topPage";
import Auth from "./component/auth";
import {Chat} from "./component/chat/chat";
import {makeStyles} from "@material-ui/core";
import {createBrowserHistory} from "history";


const app = document.getElementById('app');
const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 800,
        minWidth: 300,
        height: 600,
        backgroundColor: theme.palette.background.paper,
    },
    content: {
        width: '100%',
        maxWidth: 780,
        minWidth: 320,
        maxHeight: 600,
        minHeight: 600,
        backgroundColor: theme.palette.background.paper,
    },
}));
const history = createBrowserHistory();


ReactDOM.render(
    <div>
        <Router history={history}>
                <Switch>
                    <Route component={TopPage}/>
                    <Route component={TopPage}/>
                </Switch>
        </Router>

    </div>
    ,
    app);