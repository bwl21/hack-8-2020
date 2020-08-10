import React from 'react';
import logo from './logo.svg';
import { withStyles, WithStyles } from '@material-ui/core/styles';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    useParams,
} from "react-router-dom";
import { AssetsView } from './components/AssetsView';
import { ApplicationFrame } from './components/ApplicationFrame';

const styles = {
};

interface AppProps extends WithStyles {}

const AppImpl: React.FunctionComponent<AppProps> = (props) => {
    return <ApplicationFrame>
        <Router>
            <Switch>
                <Route path="/assets">
                    <AssetsView />
                </Route>
                <Route path="/asset/:id"></Route>
                
                <Route path="/projects"></Route>
                <Route path="/project/:id"></Route>

                <Route path="/"></Route>
            </Switch>
        </Router>
    </ApplicationFrame>
};

const App = withStyles(styles)(AppImpl);

export default App;
