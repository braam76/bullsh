local PS1 = os.getenv("USER") .. " $ "
print(PS1)
return {
	export = {
		PS1 = PS1,
	},
	aliases = {
		ll = "ls -la",
		greet = 'echo "Hello from bullsh"',
	},
}
