import * as React from 'react';
import { ThemeProvider } from 'styled-components';
import theme from '../../utils/theme';
import GlobalStyles from './GlobalStyles';
import Navbar from './navbar/Navbar';
import Footer from './footer/Footer';

interface P {
  children: React.ReactNode;
}

const Layout: React.FC<P> = ({ children }): JSX.Element => {
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyles />
      <Navbar />
      <main className="mainApp">{children}</main>
      <Footer />
    </ThemeProvider>
  );
};
export default Layout;
