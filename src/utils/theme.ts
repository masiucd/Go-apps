import { DefaultTheme } from 'styled-components';

const theme: DefaultTheme = {
  borderRadius: '1rem',
  appSize: '10px',
  shadow: {
    lightShadow: '2px 5px 3px 0px rgba(1, 2, 8, 0.5)',
    darkShadow: '4px 10px 5px 0px rgba(42, 72, 88, 0.5)',
    blackShadow: '4px 10px 5px 0px rgba(2, 2, 8, 0.5)',
  },
  app: {
    body: '#fefefe',
    text: '#1C1D21',
  },
  colors: {
    primary: '#252525',
    secondary: '#D09B69',
    danger: '#e14f3f',
    dark: '#A36F44',
    common: '#F7DEB2',
    white: '#fff',
    offWhite: '#f3f3f3',
    darkRGBA: 'rgba(0,0,0,0.4)',
  },
  size: {
    maxWidth: '1000px',
    mainSpacing: '4px',
  },
  transition: {
    mainTransition: 'all 0.3s linear',
    secondaryTransition: 'all 0.3s ease-in-out',
    quickTransition: 'all 200ms ease-in-out',
  },
};

export default theme;
