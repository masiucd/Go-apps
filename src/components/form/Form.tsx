import * as React from 'react';
import {
  StyledForm,
  FormGroup,
  StyledInput,
  StyledTextField,
  StyledLabel,
} from './FormStyles';
import { BtnPrimary } from '../styled/Buttons';

interface P {}

const Form: React.FC<P> = () => {
  return (
    <StyledForm>
      <FormGroup>
        <StyledLabel>
          <span>Name:</span>
          <StyledInput type="text" name="name" placeholder="Name" />
        </StyledLabel>
      </FormGroup>

      <FormGroup>
        <StyledLabel>
          <span>Email:</span>
          <StyledInput type="email" name="email" placeholder="Email" />
        </StyledLabel>
      </FormGroup>

      <FormGroup>
        <StyledLabel>
          <span>Message:</span>
          <StyledTextField name="message" placeholder="message" />
        </StyledLabel>
      </FormGroup>
      <FormGroup>
        <BtnPrimary type="submit">Send Message</BtnPrimary>
      </FormGroup>
    </StyledForm>
  );
};
export default Form;
