import { Link } from "react-router-dom";
import { Wrapper } from "../components/Wrapper";
import { Title } from "../components/Title";
import { Paragraph } from "../components/Paragraph";

const ForbiddenPage = () => {
  return (
    <Wrapper>
      <Title>403 — No access</Title>

      <Paragraph>
        You don't have the rights to perform this action.
      </Paragraph>

      <Paragraph>
        If you think that this is an error, please contact the administrator.
      </Paragraph>

      <Link to="/">
        Go back to the Home page
      </Link>
    </Wrapper>
  );
};

export default ForbiddenPage;
