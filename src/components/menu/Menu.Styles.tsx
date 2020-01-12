import styled from 'styled-components';

interface FooterProps {
  vegetarian: boolean;
}

export const StyledMenu = styled.div`
  section {
    margin: 3rem 0;
    h3 {
      border-bottom: 2px solid ${({ theme }) => theme.colors.secondary};
      color: ${({ theme }) => theme.colors.secondary};
    }
  }
`;
export const Grid = styled.div`
  /* display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  padding: 4rem 0;
  @media (min-width: 1000px) {
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: auto;
    grid-gap: 1em;
    justify-items: center;
  } */

  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  @media (min-width: 1000px) {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: auto;
    grid-gap: 0.5rem;
    justify-items: center;
  }
`;

export const StyledMenuItem = styled.div`
  border: 2px solid ${({ theme }) => theme.colors.primary};
  box-shadow: ${props => props.theme.shadow.lightShadow};
  position: relative;
  transition: ${({ theme }) => theme.transition.mainTransition};
  &:after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: ${({ theme }) => theme.colors.darkRGBA};
  }
  &:hover {
    box-shadow: ${props => props.theme.shadow.darkShadow};
  }

  @media (max-width: 1000px) {
    margin: 0.6em;
  }
`;

export const StyledMenuItemHead = styled.div`
  /* height: 90%; */
  .gatsby-image-wrapper {
    height: 100%;
  }
`;

export const StyledMenuItemFooter = styled.div<FooterProps>`
  color: ${({ theme, vegetarian }) =>
    vegetarian ? theme.colors.common : theme.colors.white};
  background: ${({ theme, vegetarian }) =>
    vegetarian ? theme.colors.primary : theme.colors.danger};
  padding: 0.6em 0.1em;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  span {
    z-index: 3;
    font-size: 1.6em;
    letter-spacing: 0.1rem;
    text-transform: capitalize;
  }
`;
