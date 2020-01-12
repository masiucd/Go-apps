import styled from 'styled-components';

export const StyledIntro = styled.div`
  margin: 2rem 0;
  section {
    h3 {
      border-bottom: 2px solid ${props => props.theme.colors.secondary};
      color: ${props => props.theme.colors.secondary};
    }
  }
`;

export const Text = styled.aside`
  margin: 1.4rem auto;
  padding: 0.6em;
  width: 70vw;
  p {
    letter-spacing: 0.1rem;
    font-size: 1.6em;
    line-height: 1.7em;
  }
`;
