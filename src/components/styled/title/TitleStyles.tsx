import styled from 'styled-components';

interface Props {
  flex?: boolean;
  light?: boolean;
  bgShadow?: boolean;
}

export const StyledTitle = styled.div<Props>`
  display: ${({ flex }) => (flex ? 'flex' : '')};
  align-items: ${({ flex }) => (flex ? 'center' : null)};
  justify-content: ${({ flex }) => (flex ? 'center' : null)};
  flex-direction: ${({ flex }) => (flex ? 'column' : null)};
  position: relative;
  color: ${({ theme, light }) =>
    light ? `${theme.colors.white}` : `${theme.colors.primary}`};
  /* width: 100%; */
  background: ${({ bgShadow, theme }) =>
    bgShadow ? `${theme.colors.darkRGBA}` : null};
  width: ${({ bgShadow }) => (bgShadow ? `60%` : '100%')};

  @media (max-width: 400px) {
  }
`;

export const StyledH1 = styled.h1`
  font-size: 6em;
  padding: 0.3em 0.1em 0.1em 0.1em;
  letter-spacing: 0.26rem;
  @media (max-width: 400px) {
    font-size: 4em;
  }
`;
export const StyledH3 = styled.h3`
  font-size: 5em;
  padding: 0.3em;
  letter-spacing: 0.26rem;
  @media (max-width: 400px) {
    font-size: 3em;
  }
`;
