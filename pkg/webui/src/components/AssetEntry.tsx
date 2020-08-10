import * as React from 'react';
import { withStyles } from '@material-ui/styles';
import { TableRow, Theme, createStyles, Accordion, AccordionSummary, AccordionDetails } from '@material-ui/core';
import { Asset } from '../generated/graphql';

const styles = (theme: Theme) => createStyles({
    
});

export interface AssetPreviewProps {
    open: boolean
    asset: Asset
}

const AssetPreviewImpl: React.FunctionComponent<AssetPreviewProps> = props => {
    return <Accordion expanded={props.open}>
        <AccordionSummary>Preview</AccordionSummary>
        <AccordionDetails>
            details go here
        </AccordionDetails>
    </Accordion>
}

export const AssetPreview = withStyles(styles)(AssetPreviewImpl);