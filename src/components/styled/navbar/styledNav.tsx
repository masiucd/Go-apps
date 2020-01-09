import styled from 'styled-components';
import { fadeIn } from '../../../utils/animation';

export const StyledNav = styled.nav`
  padding: 0.8em 0.1em;
  background: none;
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;
  #menuIcon {
    position: absolute;
    top: 1rem;
    right: 1rem;
    cursor: pointer;
    z-index: 8;
    transition: ${({ theme }) => theme.transition.mainTransition};
    &:hover {
      color: ${({ theme }) => theme.colors.danger};
    }
  }
  @media (min-width: 700px) {
    #menuIcon {
      display: none;
    }
  }
`;

export const AppTitle = styled.div`
  padding: 0.5em;
  margin: 0 0.3em;
  /* font-family: 'Pontano Sans', sans-serif; */
  font-family: 'Satisfy', cursive;
  display: flex;
  align-items: center;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  transition: ${({ theme }) => theme.transition.mainTransition};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  &:hover {
    box-shadow: ${({ theme }) => theme.shadow.darkShadow};
  }
  h3 {
    font-size: 3em;
    padding: 0.2em;
    letter-spacing: 0.1rem;
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
      letter-spacing: 0.1rem;

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
  z-index: 5;
      /* position: relative; */
  /* background: ${({ theme }) => theme.colors.darkRGBA}; */
  background: rgba(5,5,5,0.8);
  top: 0rem;
  width: 100vw;
  /* min-height: 220vw; */
  min-height: 100%;
  justify-content: center;
  align-items: center;
  animation: ${fadeIn} 300ms ease-in-out;
  li{
    a{
      color: ${({ theme }) => theme.colors.white};
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
