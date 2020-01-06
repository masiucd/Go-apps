import styled from 'styled-components';

export const StyledFooter = styled.footer`
  background: none;
  display: flex;
  align-items: center;
  font-family: 'Pontano Sans', sans-serif;
  @media (min-width: 700px) {
    justify-content: space-between;
  }
  @media (max-width: 700px) {
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
`;

export const FooterTitle = styled.section`
  h3 {
    font-size: 3em;
    letter-spacing: 0.1rem;
  }
  small {
    font-size: 1.1em;
    letter-spacing: 0.2rem;
  }
  @media (max-width: 700px) {
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    h3,
    small {
      text-align: center;
      display: block;
      margin: 0 auto;
      padding: 0.2em 0;
    }
  }
`;

export const Icons = styled.ul`
  display: flex;
  li {
    transition: ${({ theme }) => theme.transition.mainTransition};
    a {
      transition: ${({ theme }) => theme.transition.mainTransition};
      margin: 0 0.7em;
      &:hover {
        color: ${({ theme }) => theme.colors.danger};
      }
    }
  }
  @media (max-width: 700px) {
    padding: 0.2em 0;
    justify-content: center;
    align-items: center;

    width: 100%;
  }
`;
