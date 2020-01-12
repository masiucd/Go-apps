import styled from 'styled-components';

export const StyledMenuWrapper = styled.div`
  section {
    margin-bottom: 4.5rem;
    h1 {
    }

    h3 {
      border-bottom: 2px solid ${({ theme }) => theme.colors.secondary};
      color: ${({ theme }) => theme.colors.secondary};
    }
  }
  a {
    color: ${({ theme }) => theme.colors.secondary};
    border: 2px solid ${({ theme }) => theme.colors.primary};
  }

  margin: 1rem 0 3rem 0;
`;

export const StyledMenu = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  @media (min-width: 700px) {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: auto;
    grid-gap: 0.5rem;
    justify-items: center;
  }
`;

export const StyledMenuItem = styled.div`
  border: 1px solid ${({ theme }) => theme.colors.secondary};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  transition: ${({ theme }) => theme.transition.mainTransition};
  &:hover {
    box-shadow: ${({ theme }) => theme.shadow.darkShadow};
    transform: scale(1.02);
    .body {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }
  }
  &:after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: ${({ theme }) => theme.colors.darkRGBA};
  }

  @media (max-width: 700px) {
    margin: 0.5em;
  }
  position: relative;
`;

export const ImgWrap = styled.div`
  height: 100%;
`;

export const ItemBody = styled.div`
  display: none;
  position: absolute;
  left: 50%;
  transform: translate(-50%, 0);
  top: 34%;
  color: ${({ theme }) => theme.colors.white};
  z-index: 3;
  width: 100%;
  span {
    font-size: 1.9em;
    letter-spacing: 0.1rem;
  }
  a {
    margin-top: 0.4em;
    width: 7rem;
  }
`;
