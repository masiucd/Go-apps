import { createGlobalStyle } from 'styled-components';

export default createGlobalStyle`

/* @import url('https://fonts.googleapis.com/css?family=Barlow+Condensed:300,300i,400,400i,500,500i,600,600i,700,700i,800,800i,900,900i|Josefin+Sans:300i,400,400i,600,600i,700,700i&display=swap'); */
@import url('https://fonts.googleapis.com/css?family=Kanit:300i,400,500,500i,700,700i,800,800i,900&display=swap');
@import url('https://fonts.googleapis.com/css?family=Niramit:400,400i,500,500i,600&display=swap');
@import url('https://fonts.googleapis.com/css?family=Pontano+Sans&display=swap');
@import url('https://fonts.googleapis.com/css?family=Satisfy&display=swap');


*::before,*::after,*{
  margin: 0;
  padding: 0;
  box-sizing: inherit;
}

  body {
    box-sizing: border-box;
    background: ${props => props.theme.app.body};
    line-height: 1.4;
    color: ${({ theme }) => theme.app.text};
    height: 100%;
    display: flex;
    flex-direction: column;
    font-family: 'Kanit', sans-serif;
    font-size: 10px;
    }
  main {
    flex-grow: 1 auto;
  }

  h1,
  h2 {
    margin-bottom: 0.25rem;

  }
  p {
    margin-bottom: 1.25rem;
  }

  ul,
  li{
    list-style: none;
  }

  a {
    text-decoration: none;
    color: ${({ theme }) => theme.colors.primary};
  }

  h1 {
    font-size: 5rem;
    letter-spacing: .2rem;
  }
  h2 {
    font-size: 4rem;
  }
  h3 {
    font-size: 3.6rem;
  }
  h4 {
    font-size: 3rem;
  }
  h5 {
    font-size: 2.5rem;
  }
  p {
    font-size: 1.8rem;
  }
`;
