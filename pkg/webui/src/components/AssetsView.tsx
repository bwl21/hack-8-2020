import * as React from 'react';
import { withStyles, WithStyles, createStyles, Theme } from '@material-ui/core';

const styles = (theme: Theme) => createStyles({
    
});

export interface AssetsViewProps extends WithStyles { }

interface AssetsViewState {
    search?: string;
}

class AssetsViewImpl extends React.Component<AssetsViewProps, AssetsViewState> {

    constructor(props: AssetsViewProps) {
        super(props);

        this.state = {};
    }

    public render() {
        return <div></div>
    }

}

export const AssetsView = withStyles(styles)(AssetsViewImpl);
