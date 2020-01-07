import * as React from 'react';
import {
  StyledForm,
  FormGroup,
  StyledInput,
  StyledTextField,
} from './FormStyles';

interface P {}

const Form: React.FC<P> = () => {
  return (
    <StyledForm>
      <FormGroup>
        <StyledInput type="text" />
      </FormGroup>

      <FormGroup>
        <StyledInput type="email" />
      </FormGroup>

      <FormGroup>
        <StyledTextField></StyledTextField>
      </FormGroup>
    </StyledForm>
  );
};
export default Form;
