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
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: auto;
  grid-gap: 1em;
  justify-items: center;
  padding: 4rem 0;
`;

export const StyledMenuItem = styled.div`
  border: 2px solid ${({ theme }) => theme.colors.primary};
  box-shadow: ${props => props.theme.shadow.lightShadow};
`;
export const StyledMenuItemHead = styled.div`
  height: 90%;
  .gatsby-image-wrapper {
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
    font-size: 1.6em;
    letter-spacing: 0.1rem;
  }
`;
