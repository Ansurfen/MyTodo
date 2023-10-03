job("apk", function(ctx)
    sh("flutter build apk")
end)

job("web", function(ctx)
    sh("flutter build web")
end)

jobs("all", "apk", "web")
