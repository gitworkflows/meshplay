describe('Visit Meshplay', () => {
  it('Visits meshery', () => {
    cy.selectProviderNone();
    cy.visit('/');
  });
});

describe('Visit Meshplay settings', () => {
  it('Visits meshery settings page', () => {
    cy.selectProviderNone();
    cy.visit('/settings');
  });
});
