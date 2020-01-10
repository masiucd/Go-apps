import styled from 'styled-components';

export const StyledInfo = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  margin-top: 5rem;
  section {
    justify-content: center;
    align-items: center;
    text-align: center;
    h1 {
      font-size: 4em;
    }
    h3 {
      position: relative;
      &:after {
        content: '';
        width: 20%;
        left: 50%;
        transform: translate(-50%, 0);
        height: 2px;
        background: red;
        position: absolute;
        bottom: 0.8rem;
      }
    }
  }
`;

export const Text = styled.div`
  padding: 0.4em;
  width: 87%;
  letter-spacing: 0.2em;
  line-height: 2.2em;
  p {
    font-size: 1.6em;
    margin: 1em 0 2em 0;
  }
`;
