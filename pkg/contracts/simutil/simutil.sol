pragma solidity ^0.6;

interface IERC20 {
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    function approve(address _spender, uint256 _value) external returns (bool success);

    function transfer(address _to, uint256 _value) external returns (bool success);

    function transferFrom(
        address _from,
        address _to,
        uint256 _value
    ) external returns (bool success);

    function allowance(address _owner, address _spender) external view returns (uint256 remaining);

    function balanceOf(address _owner) external view returns (uint256 balance);

    function decimals() external view returns (uint8 digits);

    function totalSupply() external view returns (uint256 supply);
}

contract SimUtil {
    IERC20 internal constant ETH_TOKEN_ADDRESS = IERC20(
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE
    );

    function getBalances(address[] calldata users, IERC20[] calldata tokens) public view returns (uint256[]memory) {
        uint[] memory result = new uint[](users.length * tokens.length);
        for (uint i = 0; i < users.length; i++) {
            for (uint j = 0; j < tokens.length; j++) {
                if (tokens[j] == ETH_TOKEN_ADDRESS) {
                    result[i * tokens.length + j] = users[i].balance;
                } else {
                    result[i * tokens.length + j] = tokens[j].balanceOf(users[i]);
                }
            }
        }
        return result;
    }
}