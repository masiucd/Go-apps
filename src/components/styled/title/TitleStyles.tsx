import styled from 'styled-components';

interface Props {
  flex?: boolean;
}

export const StyledTitle = styled.div<Props>`
  display: ${({ flex }) => (flex ? 'flex' : '')};
  align-items: ${({ flex }) => (flex ? 'center' : null)};
  justify-content: ${({ flex }) => (flex ? 'center' : null)};
  flex-direction: ${({ flex }) => (flex ? 'column' : null)};
`;

export const StyledH1 = styled.h1`
  font-size: 6em;
  padding: 0.3em 0.1em 0.1em 0.1em;
`;
export const StyledH3 = styled.h3`
  font-size: 5em;
  padding: 0.3em;
`;
