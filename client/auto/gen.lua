job("json", function()
    sh("flutter packages pub run build_runner build")
end)

job("binding", function()

end)