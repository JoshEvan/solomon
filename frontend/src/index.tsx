import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter,Route, Switch } from "react-router-dom";
import { ItemPage } from './pages/ItemPage';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core';
import createTypography from '@material-ui/core/styles/createTypography';
import createPalette from '@material-ui/core/styles/createPalette';
import { red } from '@material-ui/core/colors';

export default function App(): JSX.Element {
    const THEME = (() => {
        const palette = createPalette({
          type: 'light',
          primary: {
            main: '#039be5',
          },
          secondary: red,
        });
      
        const typography = createTypography(palette, {
          fontFamily: '"Quicksand"',
        });
      
        return createMuiTheme({
          palette: palette,
          typography: typography,
        });
      })();
    
    return (
        <MuiThemeProvider theme={THEME}>
        <BrowserRouter>
            <Switch>
                <Route
                    path = "/" exact component={ItemPage}
                />
                <Route
                    path = "/" component={() => <div>404 - page not found</div>}
                />
            </Switch>

        </BrowserRouter>
        </MuiThemeProvider>
                
    )
}

const root = document.getElementById('app-root')

ReactDOM.render(<App/>, root)