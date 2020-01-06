import styled from 'styled-components';

export const StyledNav = styled.nav`
  padding: 0.8em 0.1em;
  background: none;
  border: 2px solid red;
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;
  #menuIcon {
    position: absolute;
    top: 1rem;
    right: 1rem;
    cursor: pointer;
    z-index: 5;
    transition: ${({ theme }) => theme.transition.mainTransition};
    &:hover {
      color: ${({ theme }) => theme.colors.dark};
    }
  }
  @media (min-width: 700px) {
    #menuIcon {
      display: none;
    }
  }
`;

export const AppTitle = styled.div`
  padding: 0.3em;
  font-family: 'Cinzel', serif;
  display: flex;
  align-items: center;
  h3 {
    font-size: 3em;
    padding: 0.2em;
  }
`;

export const StyledList = styled.ul`
  display: none;
  li {
    padding: 0 0.3em;
    transition: ${({ theme }) => theme.transition.mainTransition};
    display: flex;
    a {
      font-size: 1.5em;
      text-transform: uppercase;
      transition: ${({ theme }) => theme.transition.mainTransition};
      padding: 0.3em 0.2em;
      &:hover {
        color: ${({ theme }) => theme.colors.danger};
        box-shadow: ${({ theme }) => theme.shadow.lightShadow};
      }
    }
  }
  @media (min-width: 700px) {
    display: flex;
  }
`;

export const SmallListStyles = styled(StyledList)`
  position: fixed;
  display: flex;
  flex-direction: column;
  /* background: ${({ theme }) => theme.colors.darkRGBA}; */
  background: rgba(55,55,55,0.6);
  top: 0rem;
  width: 100vw;
  min-height: 170vw;
  justify-content: center;
  align-items: center;
  li{
    a{
      font-size: 3em;
      &:hover{
        box-shadow:none;
      }
    }
  }
  @media (min-width: 700px) {
    display: none;
  }
`;
