import styled from 'styled-components';

export const StyledFoodGrid = styled.div`
  display: grid;
  grid-template-columns: auto;
  grid-row-gap: 1rem;
  margin: 2rem 0;

  @media (min-width: 576px) {
    grid-template-columns: 1fr 1fr;
    grid-column-gap: 1rem;
  }
  @media (min-width: 768px) {
    grid-template-columns: repeat(3, 1fr);
    grid-column-gap: 1rem;
  }
  @media (min-width: 992px) {
    .gatsby-image-wrapper {
      height: 100% !important;
    }

    grid-template-areas:
      'one one two two'
      'one one three three';
  }
  .gatsby-image-wrapper {
    height: 100% !important;
  }
`;

export const GridItem = styled.div`
  position: relative;
  transition: ${({ theme }) => theme.transition.quickTransition};

  @media (min-width: 992px) {
    &:nth-child(1) {
      grid-area: one;
    }
    &:nth-child(2) {
      grid-area: two;
    }
    &:nth-child(3) {
      grid-area: three;
    }
  }

  p {
    position: absolute;
    top: 0;
    left: 0;
    color: ${({ theme }) => theme.colors.common};
    background: ${({ theme }) => theme.colors.darkRGBA};
    padding: 0.1rem 0.3rem;
    text-transform: capitalize;
    z-index: 3;
    display: none;
  }

  &:after {
    content: '';
    background: rgba(0, 0, 0, 0.3);
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
  }

  cursor: pointer;
  &:hover {
    p {
      display: block;
    }
  }
`;
