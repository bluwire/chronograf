import React, {PropTypes} from 'react'

import {MEMBER_ROLE} from 'src/auth/Authorized'

// This is a non-editable organization row, used currently for DEFAULT_ORG
const DefaultOrganization = ({organization}) =>
  <div className="orgs-table--org">
    <div className="orgs-table--id">
      {organization.id}
    </div>
    <div className="orgs-table--name-disabled">
      {organization.name}
    </div>
    <div className="orgs-table--default-role-disabled">
      {MEMBER_ROLE}
    </div>
    <button
      className="btn btn-sm btn-default btn-square orgs-table--delete"
      disabled={true}
    >
      <span className="icon trash" />
    </button>
  </div>

const {shape, string} = PropTypes

DefaultOrganization.propTypes = {
  organization: shape({
    id: string,
    name: string.isRequired,
  }).isRequired,
}

export default DefaultOrganization