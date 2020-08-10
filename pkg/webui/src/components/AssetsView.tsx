import * as React from 'react';
import { withStyles, WithStyles, createStyles, Theme } from '@material-ui/core';
import { Asset, AssetsComponent } from '../generated/graphql';

const styles = (theme: Theme) => createStyles({
    
});

export interface AssetsViewProps extends WithStyles {
}

interface AssetsViewState {
    search?: string;
    assets?: Asset[];
}

class AssetsViewImpl extends React.Component<AssetsViewProps, AssetsViewState> {

    constructor(props: AssetsViewProps) {
        super(props);

        this.state = {};
    }

    public render() {
        return <AssetsComponent>
            {result => 
                <div>
                    { result.loading && <div>loading</div> }
                    { result.data && <div>{ JSON.stringify(result.data.assets) }</div> }
                    { result.error && <div>{ result.error.message }</div> }
                </div>
            }
        </AssetsComponent>
    }

}

export const AssetsView = withStyles(styles)(AssetsViewImpl);
