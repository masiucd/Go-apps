import styled from 'styled-components';

export const PageWrapper = styled.div`
  width: 100%;
  padding-right: 1rem;
  padding-left: 1rem;
  margin-right: auto;
  margin-left: auto;
  min-height: 88vh;
  @media (min-width: 576px) {
    max-width: 540px;
  }

  @media (min-width: 768px) {
    max-width: 720px;
  }

  @media (min-width: 992px) {
    max-width: 960px;
  }

  @media (min-width: 1200px) {
    max-width: 1140px;
  }
`;

export const BtnWrapper = styled.section`
  position: absolute;
  bottom: 17rem;
  left: 50%;
  transform: translate(-50%, 0);

  width: 40rem;
  display: flex;
  justify-content: center;
`;
