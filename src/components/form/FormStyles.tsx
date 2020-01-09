import styled from 'styled-components';

export const StyledForm = styled.form`
  padding: 0.5em;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 80vw;
  margin: 4rem auto;
  background: ${({ theme }) => theme.colors.primary};
  /* margin-top: 5rem; */
`;

export const FormGroup = styled.div`
  padding: 1.3rem 0.4em;
  display: flex;
  justify-content: center;
  align-items: center;
  /* flex-direction: column; */
  width: 95%;
`;

export const StyledInput = styled.input`
  padding: 0.7em 0.3em;
  display: block;
  border-radius: 0.2em;
  width: 58vw;
  margin-top: 0.7em;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  transition: ${({ theme }) => theme.transition.mainTransition};
  outline: none;
  &:focus {
    border: ${({ theme }) => theme.colors.danger} 2px solid;
    box-shadow: ${({ theme }) => theme.shadow.darkShadow};
    width: 60vw;
    height: 2.3rem;
    span {
      display: block;
    }
  }
`;

export const StyledTextField = styled.textarea`
  padding: 0.4em;
  display: block;
  width: 58vw;
  border-radius: 0.2em;
  height: 8rem;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  box-shadow: ${({ theme }) => theme.shadow.lightShadow};
  transition: ${({ theme }) => theme.transition.mainTransition};
  margin-top: 0.7em;
  outline: none;
  &:focus {
    width: 60vw;
    box-shadow: ${({ theme }) => theme.shadow.darkShadow};
    border: ${({ theme }) => theme.colors.danger} 2px solid;
  }
`;

export const StyledLabel = styled.label`
  display: block;
  font-size: 1.5em;
  color: ${({ theme }) => theme.colors.white};
`;
