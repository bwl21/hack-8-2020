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
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';
import { AssetView, AssetViewProps } from './components/AssetView';

const styles = {
};

interface AppProps extends WithStyles {}

const AssetViewWithFilename: React.FunctionComponent<Partial<AssetViewProps>> = props => {
    const {filename} = useParams();
    return <AssetView filename={filename} {...props} />
}

const AppImpl: React.FunctionComponent<AppProps> = (props) => {
    const client = new ApolloClient({
        uri: '/graphql',
        cache: new InMemoryCache()
    });
    return <ApolloProvider client={client}>
            <ApplicationFrame>
            <Router>
                <Switch>
                    <Route path="/assets">
                        <AssetsView />
                    </Route>
                    <Route path="/asset/:filename">
                        <AssetViewWithFilename />
                    </Route>
                    
                    <Route path="/projects"></Route>
                    <Route path="/project/:id"></Route>

                    <Route path="/"></Route>
                </Switch>
            </Router>
        </ApplicationFrame>
    </ApolloProvider>
};

const App = withStyles(styles)(AppImpl);

export default App;
