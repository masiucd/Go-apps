import styled from 'styled-components';
import AniLink from 'gatsby-plugin-transition-link/AniLink';

interface BtnProps {
  position?: string;
  top?: string;
  bottom?: string;
  left?: string;
  right?: string;
  display?: string;
}

export const BtnPrimary = styled.button<BtnProps>`
  position: ${({ position }) => (position ? position : null)};
  top: ${({ top }) => (top ? top : null)};
  right: ${({ right }) => (right ? right : null)};
  left: ${({ left }) => (left ? left : null)};
  bottom: ${({ bottom }) => (bottom ? bottom : null)};
  padding: 0.2em 0.4em;
  width: 14rem;
  margin: 0 auto;
  display: block;
  background: none;
  transition: ${({ theme }) => theme.transition.mainTransition};
  color: ${({ theme }) => theme.colors.white};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  font-size: 2em;
  text-transform: uppercase;
  letter-spacing: 0.1rem;
  cursor: pointer;
  &:hover {
    color: ${({ theme }) => theme.colors.danger};
    background: ${({ theme }) => theme.colors.white};
  }
`;

export const BtnLink = styled(AniLink)<BtnProps>`
  position: ${({ position }) => (position ? position : null)};
  top: ${({ top }) => (top ? top : null)};
  right: ${({ right }) => (right ? right : null)};
  left: ${({ left }) => (left ? left : null)};
  bottom: ${({ bottom }) => (bottom ? bottom : null)};
  display: ${({ display }) => (display ? display : null)};
  text-align: center;
  padding: 0.2em 0.4em;
  width: 14rem;
  margin: 0 auto;
  background: none;
  border: ${({ theme }) => theme.colors.white} 2px solid;

  transition: ${({ theme }) => theme.transition.mainTransition};
  color: ${({ theme }) => theme.colors.white};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  font-size: 2em;
  text-transform: uppercase;
  letter-spacing: 0.1rem;

  &:hover {
    color: ${({ theme }) => theme.colors.danger};
    background: ${({ theme }) => theme.colors.white};
  }
`;
