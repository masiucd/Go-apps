import * as React from 'react';
import Form from '../form/Form';
import { StyledContactWrapper } from './Styled.Contact';

interface P {}

const Contact: React.FC<P> = () => {
  return (
    <StyledContactWrapper>
      <Form />
    </StyledContactWrapper>
  );
};
export default Contact;
