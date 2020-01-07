import styled from 'styled-components';

export const BtnPrimary = styled.button`
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
    color: ${({ theme }) => theme.colors.primary};
    background: ${({ theme }) => theme.colors.white};
  }
`;
